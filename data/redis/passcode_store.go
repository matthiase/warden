package redis

import (
	"context"
	"time"

	"github.com/matthiase/warden/util"
	"github.com/redis/go-redis/v9"
)

type PasscodeStore struct {
	*redis.Client
	Namespace string
	TTL       time.Duration
}

func (s *PasscodeStore) Create(userID string) (string, error) {
	ctx := context.TODO()
	passcode, err := util.GenerateRandomPasscode(6)
	if err != nil {
		return "", err
	}

	key := s.Namespace + passcode
	err = s.Client.Set(ctx, key, userID, s.TTL).Err()
	if err != nil {
		return "", err
	}

	return passcode, nil
}

func (s *PasscodeStore) Find(token string) (string, error) {
	ctx := context.TODO()
	key := s.Namespace + token
	userID, err := s.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return userID, nil
}

func (s *PasscodeStore) Revoke(token string) error {
	ctx := context.TODO()
	key := s.Namespace + token
	_, err := s.Client.Del(ctx, key).Result()
	if err == redis.Nil {
		return nil
	} else if err != nil {
		return err
	}

	return nil
}
