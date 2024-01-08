package internal

import (
	"log"

	"github.com/birdbox/authnz/internal/data"
)

type Application struct {
	AccountStore data.AccountStore
}

func NewApplication() (*Application, error) {
	accountStore, err := data.NewAccountStore()
	if err != nil {
		log.Fatalf("Could not set up account store: %v", err)
		return nil, err
	}

	return &Application{
		AccountStore: accountStore,
	}, nil
}
