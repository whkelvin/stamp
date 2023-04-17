package configs

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

type Configs struct {
	TestEnv                  string
	PostgresConnectionString string
	ApiKey                   string
	Port                     string
}

func (c *Configs) Init() error {
	err := godotenv.Load()
	if err != nil {
		log.Info(".env file not found, getting configs from env variables directly.")
	}

	if os.Getenv("TEST_ENV") != "" {
		c.TestEnv = os.Getenv("TEST_ENV")
	} else {
		return errors.New("Test Env String is required.")
	}

	if os.Getenv("POSTGRES_CONNECTION_STRING") != "" {
		c.PostgresConnectionString = os.Getenv("POSTGRES_CONNECTION_STRING")
	} else {
		return errors.New("Postgres Connection String is required.")
	}

	if os.Getenv("API_KEY") != "" {
		c.ApiKey = os.Getenv("API_KEY")
	} else {
		return errors.New("API key is required.")
	}

	if os.Getenv("PORT") != "" {
		c.Port = os.Getenv("PORT")
	} else {
		c.Port = ":1323"
	}
	return nil
}
