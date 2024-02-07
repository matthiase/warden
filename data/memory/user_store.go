package memory

import (
	"github.com/matthiase/warden/models"
	"github.com/segmentio/ksuid"
)

type UserStore struct {
	users map[string]*models.User
}

func NewUserStore() *UserStore {
	return &UserStore{
		users: make(map[string]*models.User),
	}
}

func (s *UserStore) Create(firstName string, lastName string, email string) (*models.User, error) {
	user := &models.User{
		ID:        ksuid.New().String(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}

	s.users[user.ID] = user
	return user, nil
}

func (s *UserStore) Find(id string) (*models.User, error) {
	user, ok := s.users[id]
	if !ok {
		return nil, nil
	}

	return user, nil
}
