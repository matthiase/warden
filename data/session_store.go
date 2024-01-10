package data

import (
	"github.com/birdbox/authnz/data/memory"
	"github.com/birdbox/authnz/models"
)

type SessionStore interface {
	Create(userID int) (models.SessionToken, error)
	//Find(t models.SessionToken) (int, error)
	//Touch(t models.SessionToken, userID int) error
	//FindAll(userID int) ([]models.SessionToken, error)
	//Revoke(t models.SessionToken) error
}

func NewSessionStore() (SessionStore, error) {
	return memory.NewSessionStore(), nil
}
