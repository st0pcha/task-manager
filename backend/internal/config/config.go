package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func Initialize() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	return &Config{
		Server: Server{
			Host:         os.Getenv("SERVER_HOST"),
			Port:         os.Getenv("SERVER_PORT"),
			AllowOrigins: strings.Split(os.Getenv("SERVER_ALLOWED_ORIGINS"), ","),
		},
	}
}
