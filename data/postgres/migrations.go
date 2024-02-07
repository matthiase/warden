package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/matthiase/warden/db"
)

var Migrations = []db.Migration{
	{
		ID: "create_users_table",
		Up: func(tx *sqlx.Tx) error {
			_, err := tx.Exec(`
				CREATE TABLE IF NOT EXISTS users (
					id TEXT PRIMARY KEY,
					first_name TEXT NOT NULL,
					last_name TEXT NOT NULL,
					email TEXT NOT NULL UNIQUE,
					locked BOOLEAN NOT NULL DEFAULT FALSE,
					created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
					deleted_at TIMESTAMPTZ
				)
			`)
			return err
		},
		Down: func(tx *sqlx.Tx) error {
			_, err := tx.Exec("DROP TABLE users")
			return err
		},
	},
}
