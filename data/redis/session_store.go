package redis

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type SessionStore struct {
	*redis.Client
	Namespace string
	TTL       time.Duration
}

func (s *SessionStore) Create(userID string) (string, error) {
	ctx := context.TODO()
	sessionID := uuid.NewString()
	sessionKey := s.Namespace + sessionID
	userKey := s.Namespace + userID
	_, err := s.Client.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.Set(ctx, sessionKey, userID, s.TTL)
		pipe.SAdd(ctx, userKey, sessionID)
		pipe.Expire(ctx, userKey, s.TTL)
		return nil
	})
	if err != nil {
		return "", err
	}
	return sessionID, nil
}

func (s *SessionStore) Find(sessionID string) (string, error) {
	ctx := context.TODO()
	key := s.Namespace + sessionID
	userID, err := s.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return userID, nil
}
