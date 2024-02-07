package redis

import (
	"net/url"

	"github.com/redis/go-redis/v9"
)

func Connect(url *url.URL) (*redis.Client, error) {
	cfg, err := redis.ParseURL(url.String())
	if err != nil {
		return nil, err
	}
	return redis.NewClient(cfg), nil
}
