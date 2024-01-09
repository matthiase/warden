package internal

import (
	"log"

	"github.com/birdbox/authnz/internal/config"
	"github.com/birdbox/authnz/internal/data"
)

type Application struct {
	UserStore data.UserStore
}

func NewApplication(cfg *config.Config) (*Application, error) {
	userStore, err := data.NewUserStore()
	if err != nil {
		log.Fatalf("Could not set up user store: %v", err)
		return nil, err
	}

	return &Application{
		UserStore: userStore,
	}, nil
}
