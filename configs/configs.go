package configs

import (
	"errors"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"os"
)

type Configs struct {
	MongoDbConnectionString    string
	MongoDbDatabaseName        string
	MongoDbPostsCollectionName string
	MongoDbUsersCollectionName string
	ApiKey                     string
	Port                       string
	JwtSecret                  string
}

func (c *Configs) Init() error {
	err := godotenv.Load()
	if err != nil {
		log.Info(".env file not found, getting configs from env variables directly.")
	}

	if os.Getenv("MONGO_DB_CONNECTION_STRING") != "" {
		c.MongoDbConnectionString = os.Getenv("MONGO_DB_CONNECTION_STRING")
	} else {
		return errors.New("Mongo db connection String is required.")
	}

	c.MongoDbDatabaseName = "stamp"
	c.MongoDbPostsCollectionName = "posts"
	c.MongoDbUsersCollectionName = "users"

	if os.Getenv("API_KEY") != "" {
		c.ApiKey = os.Getenv("API_KEY")
	} else {
		return errors.New("API key is required.")
	}

	if os.Getenv("PORT") != "" {
		c.Port = os.Getenv("PORT")
	} else {
		c.Port = "1323"
	}

	if os.Getenv("JWT_SECRET") != "" {
		c.JwtSecret = os.Getenv("JWT_SECRET")
	} else {
		return errors.New("jwt secret is required.")
	}

	return nil
}
