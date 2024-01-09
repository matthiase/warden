package memory

import (
	"github.com/birdbox/authnz/internal/models"
)

type UserStore struct {
	users map[int]*models.User
}

func NewUserStore() *UserStore {
	return &UserStore{
		users: make(map[int]*models.User),
	}
}

func (s *UserStore) Create(name string, email string) (*models.User, error) {
	user := &models.User{
		Id:    len(s.users) + 1,
		Email: email,
		Name:  name,
	}

	s.users[user.Id] = user

	return user, nil
}

func (s *UserStore) Find(id int) (*models.User, error) {
	user, ok := s.users[id]
	if !ok {
		return nil, nil
	}

	return user, nil
}
