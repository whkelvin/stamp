package config

import (
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"os"
)

type Config struct {
	PostgresConnectionString string
	ApiKey                   string
	Port                     string
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

	if os.Getenv("API_KEY") != "" {
		c.ApiKey = os.Getenv("API_KEY")
	} else {
		log.Fatal("API key is required.")
	}

	if os.Getenv("PORT") != "" {
		c.Port = os.Getenv("PORT")
	} else {
		c.Port = ":1323"
	}
	return nil
}
