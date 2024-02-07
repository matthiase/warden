package warden

import (
	"fmt"
	"time"

	"github.com/matthiase/warden/config"
	"github.com/matthiase/warden/data"
	"github.com/matthiase/warden/data/redis"
	"github.com/matthiase/warden/mailer"
)

type Application struct {
	Config        *config.Config
	PasscodeStore data.PasscodeStore
	SessionStore  data.SessionStore
	UserStore     data.UserStore
	Mailer        *mailer.Mailer
}

func NewApplication(cfg *config.Config) (*Application, error) {

	postgresClient, err := data.Connect(cfg.Database.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Postgres: %w", err)
	}

	redisClient, err := redis.Connect(cfg.Redis.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	userStore, err := data.NewUserStore(postgresClient)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize user store: %w", err)
	}

	sessionTTL := time.Duration(cfg.Session.MaxAge) * time.Second
	sessionStore, err := data.NewSessionStore(redisClient, sessionTTL)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize session store: %w", err)
	}

	passcodeTTL := time.Duration(cfg.VerificationToken.MaxAge) * time.Second
	passcodeStore, err := data.NewPasscodeStore(redisClient, passcodeTTL)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize passcode store: %w", err)
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
