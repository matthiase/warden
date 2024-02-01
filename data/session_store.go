package data

import (
	"github.com/birdbox/authnz/data/memory"
)

type SessionStore interface {
	Create(string) (string, error)
	Find(string) (string, error)
	//Touch(t models.SessionToken, userID int) error
	//FindAll(userID int) ([]models.SessionToken, error)
	//Revoke(t models.SessionToken) error
}

func NewSessionStore() (SessionStore, error) {
	return memory.NewSessionStore(), nil
}
