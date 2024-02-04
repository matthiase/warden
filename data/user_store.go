package data

import (
	"fmt"

	"github.com/birdbox/authnz/data/memory"
	"github.com/birdbox/authnz/data/postgres"
	"github.com/birdbox/authnz/models"
	"github.com/jmoiron/sqlx"
)

type UserStore interface {
	Create(firstName string, lastName string, email string) (*models.User, error)
	Find(id string) (*models.User, error)
}

func NewUserStore(db sqlx.Ext) (UserStore, error) {
	if db == nil {
		return memory.NewUserStore(), nil
	}

	switch db.DriverName() {
	case "postgres":
		return &postgres.UserStore{Ext: db}, nil
	default:
		return nil, fmt.Errorf("unsupported driver: %v", db.DriverName())
	}

}
