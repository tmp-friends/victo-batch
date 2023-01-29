package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() map[string]string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env := map[string]string{
		"BEARER_TOKEN": os.Getenv("BEARER_TOKEN"),
	}

	return env
}
