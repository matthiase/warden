package data

import (
	"github.com/birdbox/authnz/internal/data/memory"
	"github.com/birdbox/authnz/internal/models"
)

type UserStore interface {
	Create(name string, email string) (*models.User, error)
	Find(id int) (*models.User, error)
}

func NewUserStore() (UserStore, error) {
	return memory.NewUserStore(), nil
}
