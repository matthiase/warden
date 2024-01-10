package memory

import (
	"github.com/google/uuid"
)

type SessionStore struct {
	sessions map[int][]string
}

func NewSessionStore() *SessionStore {
	return &SessionStore{
		sessions: make(map[int][]string),
	}
}

func (s *SessionStore) Create(userID int) (string, error) {
	sessionID := uuid.NewString()
	s.sessions[userID] = append(s.sessions[userID], sessionID)
	return sessionID, nil
}
