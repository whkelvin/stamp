package main

import (
	//"context"
	//"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	. "github.com/whkelvin/stamp/configs"
	//. "github.com/whkelvin/stamp/pkg/api/features/write_post/controller"
	//. "github.com/whkelvin/stamp/pkg/features/write_post"
	"net/http"
)

func main() {
	configs := &Configs{}
	err := configs.Init()
	if err != nil {
		log.Fatal(err.Error())
	}

	var e *echo.Echo = echo.New()

	log.SetLevel(log.INFO)
	log.SetHeader("${time_rfc3339} ${level}")
	e.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return key == configs.ApiKey, nil
	}))

	//conn, err := pgx.Connect(context.Background(), configs.PostgresConnectionString)
	//if err != nil {
	//	log.Fatal("Postgres Connection Failed.")
	//	log.Fatal(err.Error())
	//}
	//defer conn.Close(context.Background())
	log.Debug(configs.TestEnv)

	var baseUrl = "/api/v1"

	e.GET(baseUrl+"/health", healthCheck)

	//writePostFeature := &WritePostFeature{Database: conn}
	//var writePostController *WritePostController = &WritePostController{Handler: writePostFeature.Init()}
	//writePostController.Init(baseUrl+"/post", e)

	e.Logger.Fatal(e.Start(":" + configs.Port))
}

func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Service Healthy")
}
