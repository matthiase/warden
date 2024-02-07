package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/matthiase/warden/config"
	"github.com/matthiase/warden/data"
)

func main() {
	command := os.Args[1]

	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
		os.Exit(1)
	}

	cfg := config.ReadEnv()

	switch command {
	case "upgrade":
		fmt.Println("Upgrading database")
		if err := data.Upgrade(cfg.Database.URL); err != nil {
			fmt.Printf("Error upgrading database: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Database upgraded")
	case "downgrade":
		fmt.Println("Downgrading database")
		if err := data.Downgrade(cfg.Database.URL); err != nil {
			fmt.Printf("Error downgrading database: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Database downgraded")
	default:
		fmt.Println("Unknown command")
	}
}
