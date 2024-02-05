package postgres

import (
	"github.com/birdbox/authnz/models"
	"github.com/jmoiron/sqlx"
	"github.com/segmentio/ksuid"
)

type UserStore struct {
	sqlx.Ext
}

func (db *UserStore) Find(id string) (*models.User, error) {
	sql := `
		SELECT id, first_name, last_name, email
		FROM users
		WHERE id = $1
	`

	user := models.User{}
	if err := sqlx.Get(db, &user, sql, id); err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *UserStore) Create(firstName string, lastName string, email string) (*models.User, error) {
	sql := `
		INSERT INTO users (id, first_name, last_name, email)
		VALUES ($1, $2, $3, $4)
		RETURNING id, first_name, last_name, email
	`

	uid := ksuid.New()
	result, err := db.Queryx(sql, uid.String(), firstName, lastName, email)
	if err != nil {
		return nil, err
	}

	defer result.Close()
	result.Next()

	user := models.User{}
	if err := result.StructScan(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
