package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	. "github.com/whkelvin/stamp/configs"
	. "github.com/whkelvin/stamp/pkg/api/features/get_recent_posts/controller"
	. "github.com/whkelvin/stamp/pkg/api/features/write_post/controller"
	. "github.com/whkelvin/stamp/pkg/features/get_recent_posts"
	. "github.com/whkelvin/stamp/pkg/features/write_post"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/acme/autocert"
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

	e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("https://stamp-api-dev.rootxsnowstudio.com", "https://www.stamp-api-dev.rootxsnowstudio.com")
	e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://stamp-dev.rootxsnowstudio.com", "https://www.stamp-dev.rootxsnowstudio.com", "https://stamp-dev.rootxsnowstudio.com", "https://stamp-dev.rootxsnowstudio.com"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, "x-api-key"},
	}))

	log.SetLevel(log.INFO)
	log.SetHeader("${time_rfc3339} ${level}")
	e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		Skipper:   apiKeySkipper,
		KeyLookup: "header:x-api-key",
		Validator: func(key string, c echo.Context) (bool, error) {

			return key == configs.ApiKey, nil
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

	writePostFeature := &WritePostFeature{MongoDbClient: mongoClient, MongoDbDatabaseName: configs.MongoDbDatabaseName, MongoDbCollectionName: configs.MongoDbPostsCollectionName}
	var writePostController *WritePostController = &WritePostController{Handler: writePostFeature.Init()}
	writePostController.Init(BASE_URL+"/post", e)

	getRecentPostsFeature := &GetRecentPostsFeature{MongoDbClient: mongoClient, MongoDbDatabaseName: configs.MongoDbDatabaseName, MongoDbCollectionName: configs.MongoDbPostsCollectionName}
	var getRecentPostsController *GetRecentPostsController = &GetRecentPostsController{Handler: getRecentPostsFeature.Init()}
	getRecentPostsController.Init(BASE_URL+"/posts", e)

	e.Logger.Fatal(e.StartAutoTLS(":" + configs.Port))
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
