package config

import (
	"log"
	"net/url"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Database struct {
		Url *url.URL
	}
	Server struct {
		Host   string
		Port   int
		Secret string
	}
	SessionToken struct {
		Secret string
		MaxAge int
	}
	AccessToken struct {
		Secret string
		MaxAge int
	}
}

func ReadEnv() *Config {
	cfg := &Config{}
	cfg.Database.Url = parseURL("DATABASE_URL")
	cfg.Server.Host = lookupString("SERVER_HOST", "localhost")
	cfg.Server.Port = lookupInt("SERVER_PORT", 5000)
	cfg.SessionToken.Secret = os.Getenv("SESSION_TOKEN_SECRET")
	cfg.SessionToken.MaxAge = lookupInt("SESSION_TOKEN_MAX_AGE", 86400)
	cfg.AccessToken.Secret = os.Getenv("ACCESS_TOKEN_SECRET")
	cfg.AccessToken.MaxAge = lookupInt("ACCESS_TOKEN_MAX_AGE", 3600)
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
