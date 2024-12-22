package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Initialize() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	return &Config{
		Mode: os.Getenv("MODE"),
		Server: Server{
			Host:         os.Getenv("SERVER_HOST"),
			Port:         os.Getenv("SERVER_PORT"),
			AllowOrigins: os.Getenv("SERVER_ALLOWED_ORIGINS"),
		},
		Postgres: PostgresDatabase{
			DSN: os.Getenv("POSTGRES_DSN"),
		},
	}
}
