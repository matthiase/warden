package data

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/matthiase/warden/data/memory"
	"github.com/matthiase/warden/data/postgres"
	"github.com/matthiase/warden/models"
)

type UserStore interface {
	Create(firstName string, lastName string, email string) (*models.User, error)
	Find(id string) (*models.User, error)
}

func NewUserStore(client sqlx.Ext) (UserStore, error) {
	if client == nil {
		return memory.NewUserStore(), nil
	}

	switch client.DriverName() {
	case "postgres":
		return &postgres.UserStore{Ext: client}, nil
	default:
		return nil, fmt.Errorf("unsupported driver: %v", client.DriverName())
	}
}
