package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	MongoDbConnectionString string
}

func (c *Config) Init() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	if os.Getenv("MONGO_DB_CONNECTION_STRING") != "" {
		c.MongoDbConnectionString = os.Getenv("MONGO_DB_CONNECTION_STRING")
	} else {
		c.MongoDbConnectionString = "mongodb://foo:bar@localhost:27017"
	}
	return nil
}
