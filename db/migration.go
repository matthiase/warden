package db

import "github.com/jmoiron/sqlx"

type Migration struct {
	ID   string
	Up   func(*sqlx.Tx) error
	Down func(*sqlx.Tx) error
}

/*
func Migrate(db *sqlx.DB, migrations []Migration) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, m := range migrations {
		if _, err := tx.Exec("INSERT INTO migrations (id) VALUES ($1)", m.ID); err != nil {
			return err
		}
		if err := m.Up(tx); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func Rollback(db *sqlx.DB, migrations []Migration) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for i := len(migrations) - 1; i >= 0; i-- {
		if _, err := tx.Exec("DELETE FROM migrations WHERE id = $1", migrations[i].ID); err != nil {
			return err
		}
		if err := migrations[i].Down(tx); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func EnsureMigrationsTable(db *sqlx.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS migrations (id TEXT PRIMARY KEY)")
	return err
}

func GetAppliedMigrations(db *sqlx.DB) ([]string, error) {
	rows, err := db.Query("SELECT id FROM migrations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	return ids, nil
}

func GetPendingMigrations(db *sqlx.DB, migrations []Migration) ([]Migration, error) {
	applied, err := GetAppliedMigrations(db)
	if err != nil {
		return nil, err
	}

	var pending []Migration
	for _, m := range migrations {
		found := false
		for _, a := range applied {
			if m.ID == a {
				found = true
				break
			}
		}
		if !found {
			pending = append(pending, m)
		}
	}

	return pending, nil
}

func GetLastMigrationID(migrations []Migration) string {
	if len(migrations) == 0 {
		return ""
	}
	return migrations[len(migrations)-1].ID
}

func GetMigrationByID(migrations []Migration, id string) *Migration {
	for _, m := range migrations {
		if m.ID == id {
			return &m
		}
	}
	return nil
}

func GetMigrationIDs(migrations []Migration) []string {
	ids := make([]string, len(migrations))
	for i, m := range migrations {
		ids[i] = m.ID
	}
	return ids
}
*/
