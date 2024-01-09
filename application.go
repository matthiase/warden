package authnz

import (
	"log"

	"github.com/birdbox/authnz/config"
	"github.com/birdbox/authnz/data"
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
