package main

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"
	. "github.com/whkelvin/stamp/configs"
	. "github.com/whkelvin/stamp/pkg/api/features/get_recent_posts/controller"
	. "github.com/whkelvin/stamp/pkg/api/features/write_post/controller"
	. "github.com/whkelvin/stamp/pkg/features/get_recent_posts"
	. "github.com/whkelvin/stamp/pkg/features/write_post"
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
		AllowOrigins: []string{"*"},
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

	pgxConfig, err := pgxpool.ParseConfig(configs.PostgresConnectionString)
	if err != nil {
		panic(err)
	}

	pgxConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxUUID.Register(conn.TypeMap())
		return nil
	}

	pgxConnPool, err := pgxpool.NewWithConfig(context.Background(), pgxConfig)
	if err != nil {
		log.Error("Postgres Connection Failed.")
		log.Error(err.Error())
		panic(err)
	}
	defer pgxConnPool.Close()

	e.GET(BASE_URL+"/health", healthCheck)

	writePostFeature := &WritePostFeature{ConnPool: pgxConnPool}
	var writePostController *WritePostController = &WritePostController{Handler: writePostFeature.Init()}
	writePostController.Init(BASE_URL+"/post", e)

	getRecentPostsFeature := &GetRecentPostsFeature{ConnPool: pgxConnPool}
	var getRecentPostsController *GetRecentPostsController = &GetRecentPostsController{Handler: getRecentPostsFeature.Init()}
	getRecentPostsController.Init(BASE_URL+"/posts", e)

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
