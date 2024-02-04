package authnz

import (
	"time"

	"github.com/birdbox/authnz/config"
	"github.com/birdbox/authnz/data"
	"github.com/birdbox/authnz/mailer"
)

type Application struct {
	Config        *config.Config
	PasscodeStore data.PasscodeStore
	SessionStore  data.SessionStore
	UserStore     data.UserStore
	Mailer        *mailer.Mailer
}

func NewApplication(cfg *config.Config) (*Application, error) {

	db, err := data.Connect(cfg.Database.URL)
	if err != nil {
		panic(err)
	}

	userStore, err := data.NewUserStore(db)
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

	mailer := mailer.NewMailer(&mailer.MailerConfig{
		Host:     cfg.Mailer.Host,
		Port:     cfg.Mailer.Port,
		Username: cfg.Mailer.Username,
		Password: cfg.Mailer.Password,
		Sender:   cfg.Mailer.Sender,
		Timeout:  time.Duration(cfg.Mailer.Timeout) * time.Second,
	})

	return &Application{
		Config:        cfg,
		PasscodeStore: passcodeStore,
		SessionStore:  sessionStore,
		UserStore:     userStore,
		Mailer:        mailer,
	}, nil
}
