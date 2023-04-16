package config

import (
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"os"
)

type Config struct {
	PostgresConnectionString string
}

func (c *Config) Init() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	if os.Getenv("POSTGRES_CONNECTION_STRING") != "" {
		c.PostgresConnectionString = os.Getenv("POSTGRES_CONNECTION_STRING")
	} else {
		log.Fatal("Postgres Connection String is required.")
	}
	return nil
}
