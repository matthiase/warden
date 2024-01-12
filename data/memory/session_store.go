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

func (s *SessionStore) Find(sessionID string) (int, error) {
	for userID, sessionIDs := range s.sessions {
		for _, id := range sessionIDs {
			if id == sessionID {
				return userID, nil
			}
		}
	}
	return 0, nil
}
