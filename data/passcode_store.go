package data

import (
	"github.com/birdbox/authnz/data/memory"
)

type PasscodeStore interface {
	Create(userID int) (string, error)
	Find(t string) (int, error)
	Revoke(t string) error
}

func NewPasscodeStore() (PasscodeStore, error) {
	return memory.NewPasscodeStore(), nil
}
