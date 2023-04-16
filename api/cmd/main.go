package main

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	. "github.com/whkelvin/stamp/api/config"
	. "github.com/whkelvin/stamp/api/features/write_post/controller"
	. "github.com/whkelvin/stamp/features/pkg/write_post"
)

func main() {
	var e *echo.Echo = echo.New()

	log.SetLevel(log.DEBUG)
	log.SetHeader("${time_rfc3339} ${level}")

	config := &Config{}
	config.Init()

	conn, err := pgx.Connect(context.Background(), config.PostgresConnectionString)
	if err != nil {
		log.Fatal("Postgres Connection Failed.")
		log.Fatal(err.Error())
	}
	defer conn.Close(context.Background())

	writePostFeature := &WritePostFeature{Database: conn}
	var writePostController *WritePostController = &WritePostController{Handler: writePostFeature.Init()}
	writePostController.Init("/post", e)

	e.Logger.Fatal(e.Start(":1323"))
}
