package data

import (
	"fmt"
	"net/url"

	"github.com/jmoiron/sqlx"
	"github.com/matthiase/warden/data/postgres"
)

func Connect(url *url.URL) (*sqlx.DB, error) {
	switch url.Scheme {
	case "memory":
		return nil, nil
	case "postgresql", "postgres":
		return postgres.Connect(url)
	default:
		return nil, fmt.Errorf("unsupported database: %s", url.Scheme)
	}
}

func Upgrade(url *url.URL) error {
	switch url.Scheme {
	case "memory":
		return nil
	case "postgresql", "postgres":
		return postgres.Upgrade(url)
	default:
		return fmt.Errorf("unsupported database: %s", url.Scheme)
	}
}

func Downgrade(url *url.URL) error {
	switch url.Scheme {
	case "memory":
		return nil
	case "postgresql", "postgres":
		return postgres.Downgrade(url)
	default:
		return fmt.Errorf("unsupported database: %s", url.Scheme)
	}
}
