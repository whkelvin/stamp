package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	. "github.com/whkelvin/stamp/configs"
	. "github.com/whkelvin/stamp/pkg/api/controllers"
	. "github.com/whkelvin/stamp/pkg/api/generated"
	. "github.com/whkelvin/stamp/pkg/features/get_recent_posts"
	. "github.com/whkelvin/stamp/pkg/features/write_post"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

var BASE_URL = "/api/v1"

func main() {
	configs := &Configs{}
	err := configs.Init()
	if err != nil {
		log.Error(err.Error())
	}

	var e *echo.Echo = echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://stamp-dev.rootxsnowstudio.com", "https://www.stamp-dev.rootxsnowstudio.com", "http://localhost:5173"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, "x-api-key"},
	}))

	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogLevel:  log.ERROR,
	}))

	log.SetLevel(log.INFO)
	log.SetHeader("${time_rfc3339} ${level}")
	e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		Skipper:   apiKeySkipper,
		KeyLookup: "header:x-api-key",
		Validator: func(key string, c echo.Context) (bool, error) {
			return key == configs.ApiKey, nil
		},
		ErrorHandler: func(err error, c echo.Context) error {
			return c.String(http.StatusUnauthorized, "")
		},
	}))

	ctx := context.Background()
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(configs.MongoDbConnectionString))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	e.GET(BASE_URL+"/health", healthCheck)

	writePostFeature := WritePostFeature{MongoDbClient: mongoClient, MongoDbDatabaseName: configs.MongoDbDatabaseName, MongoDbCollectionName: configs.MongoDbPostsCollectionName}
	getRecentPostsFeature := GetRecentPostsFeature{MongoDbClient: mongoClient, MongoDbDatabaseName: configs.MongoDbDatabaseName, MongoDbCollectionName: configs.MongoDbPostsCollectionName}

	var apiServer ApiServer = ApiServer{
		WritePostHandler:      writePostFeature.Init(),
		GetRecentPostsHandler: getRecentPostsFeature.Init(),
	}
	RegisterHandlersWithBaseURL(e, &apiServer, BASE_URL)

	e.Logger.Fatal(e.Start(":" + configs.Port))
}

func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Service Healthy")
}

func apiKeySkipper(c echo.Context) bool {
	if c.Path() == BASE_URL+"/posts" && c.Request().Method == "GET" {
		return true
	}
	return false
}
