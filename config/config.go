package config

import (
	"log"
	"net/url"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Application string
	Environment string
	Server      struct {
		Host   string
		Port   int
		Secret string
	}
	Database struct {
		Url *url.URL
	}
	Session struct {
		Name   string
		Secure bool
		MaxAge int
	}
	IdentityToken struct {
		MaxAge int
	}
	VerificationToken struct {
		MaxAge int
	}
	Mailer struct {
		Host     string
		Port     int
		Username string
		Password string
		Sender   string
		Timeout  int
	}
}

func ReadEnv() *Config {
	cfg := &Config{}
	cfg.Application = lookupString("APPLICATION_NAME", "Authnz")
	cfg.Environment = lookupString("ENVIRONMENT", "development")
	cfg.Server.Host = lookupString("SERVER_HOST", "localhost")
	cfg.Server.Port = lookupInt("SERVER_PORT", 5000)
	cfg.Server.Secret = os.Getenv("SERVER_SECRET")
	cfg.Database.Url = parseURL("DATABASE_URL")
	cfg.Session.Name = lookupString("SESSION_COOKIE_NAME", "authnz")
	cfg.Session.Secure = lookupBool("SESSION_COOKIE_SECURE", false)
	cfg.Session.MaxAge = lookupInt("SESSION_TOKEN_MAX_AGE", 86400)
	cfg.IdentityToken.MaxAge = lookupInt("IDENTITY_TOKEN_MAX_AGE", 3600)
	cfg.VerificationToken.MaxAge = lookupInt("VERIFICATION_TOKEN_MAX_AGE", 300)
	cfg.Mailer.Host = lookupString("SMTP_HOST", "localhost")
	cfg.Mailer.Port = lookupInt("SMTP_PORT", 1025)
	cfg.Mailer.Username = os.Getenv("SMTP_USERNAME")
	cfg.Mailer.Password = os.Getenv("SMTP_PASSWORD")
	cfg.Mailer.Sender = lookupString("SMTP_SENDER", "no-reply@localhost")
	cfg.Mailer.Timeout = lookupInt("SMTP_TIMEOUT", 5)
	return cfg
}

func parseURL(name string) *url.URL {
	if str, ok := os.LookupEnv(name); ok {
		url, err := url.ParseRequestURI(str)
		if err != nil {
			log.Fatalf("Invalid %s", name)
		}
		return url
	}
	return nil
}

func lookupInt(name string, defaultValue int) int {
	if str, ok := os.LookupEnv(name); ok {
		value, err := strconv.Atoi(str)
		if err != nil {
			log.Fatalf("Invalid %s", name)
		}
		return value
	}
	return defaultValue
}

func lookupString(name string, defaultValue string) string {
	if str, ok := os.LookupEnv(name); ok {
		return str
	}
	return defaultValue
}

func lookupBool(name string, defaultValue bool) bool {
	if str, ok := os.LookupEnv(name); ok {
		value, err := strconv.ParseBool(str)
		if err != nil {
			log.Fatalf("Invalid %s", name)
		}
		return value
	}
	return defaultValue
}
