package memory

import (
	"github.com/birdbox/authnz/internal/models"
)

type AccountStore struct {
	accounts map[int]*models.Account
}

func NewAccountStore() *AccountStore {
	return &AccountStore{
		accounts: make(map[int]*models.Account),
	}
}

func (s *AccountStore) Create(name string, email string) (*models.Account, error) {
	account := &models.Account{
		Id:    len(s.accounts) + 1,
		Email: email,
		Name:  name,
	}

	s.accounts[account.Id] = account

	return account, nil
}

func (s *AccountStore) Find(id int) (*models.Account, error) {
	account, ok := s.accounts[id]
	if !ok {
		return nil, nil
	}

	return account, nil
}
