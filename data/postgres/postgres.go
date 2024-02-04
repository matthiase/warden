package postgres

import (
	"fmt"
	"log"
	"net/url"
	"slices"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect(url *url.URL) (*sqlx.DB, error) {
	return sqlx.Connect("postgres", url.String())
}

func Upgrade(url *url.URL) error {
	db, err := Connect(url)
	if err != nil {
		return err
	}

	defer db.Close()

	if err := ensureMigrationsTable(db); err != nil {
		return err
	}

	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Get the list of migrations that have already been run
	var applied []string
	if err := tx.Select(&applied, "SELECT id FROM migrations"); err != nil {
		return fmt.Errorf("getting all migrations: %w", err)
	}

	for _, m := range Migrations {
		if slices.Contains(applied, m.ID) {
			continue
		}

		if _, err := tx.Exec("INSERT INTO migrations (id) VALUES ($1)", m.ID); err != nil {
			return fmt.Errorf("inserting migration %s: %w", m.ID, err)
		}

		if err := m.Up(tx); err != nil {
			return err
		}

		log.Printf("Applied migration %s", m.ID)
	}

	return tx.Commit()
}

func Downgrade(url *url.URL) error {
	db, err := Connect(url)
	if err != nil {
		return err
	}

	defer db.Close()

	if err := ensureMigrationsTable(db); err != nil {
		return err
	}

	// Retrieve a list of all migrations that have been run
	var migrations []string
	if err := db.Select(&migrations, "SELECT id FROM migrations"); err != nil {
		return fmt.Errorf("getting all migrations: %w", err)
	}

	// If there are no migrations, we don't need to do anything
	if len(migrations) == 0 {
		return nil
	}

	// Get the most recent migration from the list
	lastMigrationId := migrations[len(migrations)-1]

	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, m := range Migrations {
		if m.ID == lastMigrationId {
			log.Printf("Reverting migration %s", m.ID)
			if err := m.Down(tx); err != nil {
				return err
			}
		}
	}

	if _, err := tx.Exec("DELETE FROM migrations WHERE id = $1", lastMigrationId); err != nil {
		return fmt.Errorf("deleting migration: %w", err)
	}

	return tx.Commit()
}

func ensureMigrationsTable(db *sqlx.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS migrations (
			id TEXT PRIMARY KEY
		)
	`)
	return err
}
