package memory

import (
	"github.com/birdbox/authnz/util"
)

type PasscodeStore struct {
	passcodes map[string]string
}

func NewPasscodeStore() *PasscodeStore {
	return &PasscodeStore{
		passcodes: make(map[string]string),
	}
}

func (s *PasscodeStore) Create(userID string) (string, error) {
	passcode, err := util.GenerateRandomPasscode(6)
	if err != nil {
		return "", err
	}
	s.passcodes[passcode] = userID
	return string(passcode), nil
}

func (s *PasscodeStore) Find(t string) (string, error) {
	return s.passcodes[t], nil
}

func (s *PasscodeStore) Revoke(t string) error {
	delete(s.passcodes, t)
	return nil
}
