package memory

import (
	"fmt"

	"github.com/google/uuid"
)

type SessionStore struct {
	sessions map[string][]string
}

func NewSessionStore() *SessionStore {
	return &SessionStore{
		sessions: make(map[string][]string),
	}
}

func (s *SessionStore) Create(userID string) (string, error) {
	sessionID := uuid.NewString()
	s.sessions[userID] = append(s.sessions[userID], sessionID)
	return sessionID, nil
}

func (s *SessionStore) Find(sessionID string) (string, error) {
	for userID, sessionIDs := range s.sessions {
		for _, id := range sessionIDs {
			if id == sessionID {
				return userID, nil
			}
		}
	}
	return "", fmt.Errorf("session not found")
}
