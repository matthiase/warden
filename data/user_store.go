package data

import (
	"github.com/birdbox/authnz/data/memory"
	"github.com/birdbox/authnz/models"
)

type UserStore interface {
	Create(firstName string, lastName string, email string) (*models.User, error)
	Find(id string) (*models.User, error)
}

func NewUserStore() (UserStore, error) {
	return memory.NewUserStore(), nil
}
