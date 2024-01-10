package memory

import (
	"encoding/hex"

	"github.com/birdbox/authnz/models"
	"github.com/birdbox/authnz/util"
)

type SessionStore struct {
	sessions map[int][]models.SessionToken
}

func NewSessionStore() *SessionStore {
	return &SessionStore{
		sessions: make(map[int][]models.SessionToken),
	}
}

func (s *SessionStore) Create(userID int) (models.SessionToken, error) {
	binToken, err := util.RandomToken()
	if err != nil {
		return "", err
	}
	token := models.SessionToken(hex.EncodeToString(binToken))
	s.sessions[userID] = append(s.sessions[userID], token)
	return token, nil
}

//func (s *sessionStore) Find(t models.SessionToken) (int, error) {
//	return s.accountByToken[t], nil
//}
//
//func (s *sessionStore) Touch(t models.SessionToken, accountID int) error {
//	return nil
//}
//
//func (s *sessionStore) FindAll(accountID int) ([]models.SessionToken, error) {
//	return s.sessions[accountID], nil
//}
//
//func (s *sessionStore) Revoke(t models.SessionToken) error {
//	accountID := s.accountByToken[t]
//	if accountID != 0 {
//		delete(s.accountByToken, t)
//		s.sessions[accountID] = without(t, s.sessions[accountID])
//	}
//	return nil
//}
//
//func without(needle models.SessionToken, haystack []models.SessionToken) []models.SessionToken {
//	for idx, elem := range haystack {
//		if elem == needle {
//			return append(haystack[:idx], haystack[idx+1:]...)
//		}
//	}
//	return haystack
//}
