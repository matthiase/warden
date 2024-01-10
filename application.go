package authnz

import (
	"github.com/birdbox/authnz/config"
	"github.com/birdbox/authnz/data"
)

type Application struct {
	Config        *config.Config
	PasscodeStore data.PasscodeStore
	SessionStore  data.SessionStore
	UserStore     data.UserStore
}

func NewApplication(cfg *config.Config) (*Application, error) {
	userStore, err := data.NewUserStore()
	if err != nil {
		return nil, err
	}

	sessionStore, err := data.NewSessionStore()
	if err != nil {
		return nil, err
	}

	passcodeStore, err := data.NewPasscodeStore()
	if err != nil {
		return nil, err
	}

	return &Application{
		Config:        cfg,
		PasscodeStore: passcodeStore,
		SessionStore:  sessionStore,
		UserStore:     userStore,
	}, nil
}
