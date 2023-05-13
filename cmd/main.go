package main

import (
	"context"
	"net/http"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	. "github.com/whkelvin/stamp/configs"
	. "github.com/whkelvin/stamp/pkg/api/controllers"
	. "github.com/whkelvin/stamp/pkg/api/generated"
	. "github.com/whkelvin/stamp/pkg/features/get_recent_posts"
	. "github.com/whkelvin/stamp/pkg/features/log_in"
	. "github.com/whkelvin/stamp/pkg/features/refresh_token"
	. "github.com/whkelvin/stamp/pkg/features/write_post"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var BASE_URL = "/api/v1"

func main() {
	configs := getConfigs()

	var e *echo.Echo = echo.New()
	setupMiddleWare(e, configs)

	e.GET(BASE_URL+"/health", healthCheck)

	ctx := context.Background()
	mongoClient := getMongoClient(ctx, configs)
	defer func() {
		if err := mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	initFeatures(configs, e, mongoClient)
	e.Logger.Fatal(e.Start(":" + configs.Port))
}

func getConfigs() *Configs {
	configs := &Configs{}
	err := configs.Init()
	if err != nil {
		panic(err.Error())
	}
	return configs
}

func getMongoClient(ctx context.Context, configs *Configs) *mongo.Client {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(configs.MongoDbConnectionString))
	if err != nil {
		panic(err.Error())
	}
	return client
}

func setupMiddleWare(e *echo.Echo, configs *Configs) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"https://stamp-dev.rootxsnowstudio.com",
			"https://www.stamp-dev.rootxsnowstudio.com",
			"http://localhost:5173", // TODO: disable this in prod
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
		},
	}))

	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogLevel:  log.ERROR,
	}))

	log.SetLevel(log.INFO)
	log.SetHeader("${time_rfc3339} ${level}")

	e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(configs.JwtSecret),
		Skipper:    jwtAuthSkipper,
	}))

	// limit the application to 20 requests/sec using the default in-memory
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
}

func initFeatures(configs *Configs, e *echo.Echo, mongoClient *mongo.Client) {

	writePostFeature := WritePostFeature{MongoDbClient: mongoClient, MongoDbDatabaseName: configs.MongoDbDatabaseName, MongoDbCollectionName: configs.MongoDbPostsCollectionName}

	getRecentPostsFeature := GetRecentPostsFeature{MongoDbClient: mongoClient, MongoDbDatabaseName: configs.MongoDbDatabaseName, MongoDbCollectionName: configs.MongoDbPostsCollectionName}

	logInFeature := LogInFeature{JwtSecret: configs.JwtSecret, MongoDbClient: mongoClient, MongoDbDatabaseName: configs.MongoDbDatabaseName, MongoDbCollectionName: configs.MongoDbUsersCollectionName}

	refreshTokenFeature := RefreshTokenFeature{JwtSecret: configs.JwtSecret}

	var apiServer ApiServer = ApiServer{
		WritePostHandler:      writePostFeature.Init(),
		GetRecentPostsHandler: getRecentPostsFeature.Init(),
		LogInHandler:          logInFeature.Init(),
		RefreshTokenHandler:   refreshTokenFeature.Init(),
	}
	RegisterHandlersWithBaseURL(e, &apiServer, BASE_URL)
}

func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Service Healthy")
}

func jwtAuthSkipper(c echo.Context) bool {
	if c.Request().Method == "GET" && c.Path() == BASE_URL+"/posts" {
		return true
	}

	if c.Request().Method == "POST" && c.Path() == BASE_URL+"/login" {
		return true
	}
	return false
}
