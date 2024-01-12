package data

import (
	"github.com/birdbox/authnz/data/memory"
)

type SessionStore interface {
	Create(int) (string, error)
	Find(string) (int, error)
	//Touch(t models.SessionToken, userID int) error
	//FindAll(userID int) ([]models.SessionToken, error)
	//Revoke(t models.SessionToken) error
}

func NewSessionStore() (SessionStore, error) {
	return memory.NewSessionStore(), nil
}
