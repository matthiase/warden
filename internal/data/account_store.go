package data

import (
	"github.com/birdbox/authnz/internal/data/memory"
	"github.com/birdbox/authnz/internal/models"
)

type AccountStore interface {
	Create(name string, email string) (*models.Account, error)
	Find(id int) (*models.Account, error)
}

func NewAccountStore() (AccountStore, error) {
	return memory.NewAccountStore(), nil
}
