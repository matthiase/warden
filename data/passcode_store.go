package data

import (
	"github.com/birdbox/authnz/data/memory"
	"github.com/birdbox/authnz/models"
)

type PasscodeStore interface {
	Create(userID int) (models.Passcode, error)
	Find(t models.Passcode) (int, error)
	Revoke(t models.Passcode) error
}

func NewPasscodeStore() (PasscodeStore, error) {
	return memory.NewPasscodeStore(), nil
}
