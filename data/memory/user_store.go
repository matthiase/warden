package memory

import (
	"github.com/birdbox/authnz/models"
)

type UserStore struct {
	users map[int]*models.User
}

func NewUserStore() *UserStore {
	return &UserStore{
		users: make(map[int]*models.User),
	}
}

func (s *UserStore) Create(firstName string, lastName string, email string) (*models.User, error) {
	user := &models.User{
		ID:        len(s.users) + 1,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}

	s.users[user.ID] = user

	return user, nil
}

func (s *UserStore) Find(id int) (*models.User, error) {
	user, ok := s.users[id]
	if !ok {
		return nil, nil
	}

	return user, nil
}
