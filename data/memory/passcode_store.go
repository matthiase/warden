package memory

import (
	"github.com/birdbox/authnz/util"
)

type PasscodeStore struct {
	passcodes map[string]int
}

func NewPasscodeStore() *PasscodeStore {
	return &PasscodeStore{
		passcodes: make(map[string]int),
	}
}

func (s *PasscodeStore) Create(userID int) (string, error) {
	passcode, err := util.GenerateRandomPasscode(6)
	if err != nil {
		return "", err
	}
	s.passcodes[passcode] = userID
	return string(passcode), nil
}

func (s *PasscodeStore) Find(t string) (int, error) {
	return s.passcodes[t], nil
}

func (s *PasscodeStore) Revoke(t string) error {
	delete(s.passcodes, t)
	return nil
}
