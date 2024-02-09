package data

import (
	"time"

	"github.com/matthiase/warden/data/memory"
	"github.com/matthiase/warden/data/redis"
	db "github.com/redis/go-redis/v9"
)

type PasscodeStore interface {
	Create(string) (string, error)
	Find(string) (string, error)
	Revoke(string) error
}

func NewPasscodeStore(client *db.Client, ttl time.Duration) (PasscodeStore, error) {
	if client == nil {
		return memory.NewPasscodeStore(), nil
	}

	return &redis.PasscodeStore{
		Client:    client,
		Namespace: "warden:passcode:",
		TTL:       ttl,
	}, nil
}
