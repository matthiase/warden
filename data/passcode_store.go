package data

import (
	"github.com/birdbox/authnz/data/memory"
)

type PasscodeStore interface {
	Create(userID string) (string, error)
	Find(t string) (string, error)
	Revoke(t string) error
}

func NewPasscodeStore() (PasscodeStore, error) {
	return memory.NewPasscodeStore(), nil
}
