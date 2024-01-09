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
		Port int
	}
}

func ReadEnv() *Config {
	cfg := &Config{}
	cfg.Database.Url = lookupURL("DATABASE_URL")
	cfg.Server.Port = lookupInt("SERVER_PORT", 5000)
	return cfg
}

func lookupURL(name string) *url.URL {
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
