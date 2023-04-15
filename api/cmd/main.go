package main

import (
	//"context"
	//"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	//_ "github.com/mattn/go-sqlite3"
	. "github.com/whkelvin/stamp/api/features/write_post/controller"
	. "github.com/whkelvin/stamp/features/pkg/write_post/handler"
)

func main() {
	var e *echo.Echo = echo.New()

	log.SetLevel(log.DEBUG)
	log.SetHeader("${time_rfc3339} ${level}")

	//db, err := sql.Open("sqlite3", "/home/whkelvin/Projects/golang/stamp/database/prisma/stamp.db")

	//if err != nil {
	//	log.Fatal("Opening db failed")
	//}

	//var dbService *service.DbService = &service.DbService{Ctx: context.Background(), Db: db}

	//var getUserDbService *GetUserDbService = &GetUserDbService{DbService: dbService}
	//var getUserHandler *GetUserHandler = &GetUserHandler{GetUserDbService: getUserDbService}
	var writePostHandler *WritePostHandler = &WritePostHandler{}
	var writePostController *WritePostController = &WritePostController{Handler: writePostHandler}

	writePostController.Init("/post", e)

	e.Logger.Fatal(e.Start(":1323"))
}
