package memory

import (
	"github.com/birdbox/authnz/models"
	"github.com/birdbox/authnz/util"
)

type PasscodeStore struct {
	passcodes map[models.Passcode]int
}

func NewPasscodeStore() *PasscodeStore {
	return &PasscodeStore{
		passcodes: make(map[models.Passcode]int),
	}
}

func (s *PasscodeStore) Create(userID int) (models.Passcode, error) {
	hexToken, err := util.GenerateRandomPasscode(6)
	if err != nil {
		return "", err
	}
	passcode := models.Passcode(hexToken)
	s.passcodes[passcode] = userID
	return models.Passcode(passcode), nil
}

func (s *PasscodeStore) Find(t models.Passcode) (int, error) {
	return s.passcodes[t], nil
}

func (s *PasscodeStore) Revoke(t models.Passcode) error {
	delete(s.passcodes, t)
	return nil
}
