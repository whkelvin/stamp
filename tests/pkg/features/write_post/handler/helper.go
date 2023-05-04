package handler

import (
	dbModels "github.com/whkelvin/stamp/pkg/features/write_post/db/models"
	handlerModels "github.com/whkelvin/stamp/pkg/features/write_post/handler/models"
	"time"
)

var id string = "6451167307ff6dafe2ea09ea"
var link string = "https://youtu.be/c4OyfL5o7DU"
var rootDomain string = "youtube.com"
var title string = "test title"
var description string = "test description"
var createdDate string = time.Now().UTC().Format(time.RFC3339)

func GetHandlerRequest() handlerModels.Request {
	return handlerModels.Request{
		Link:        link,
		RootDomain:  rootDomain,
		Title:       title,
		Description: description,
	}
}

func GetHandlerResponse() handlerModels.Response {
	return handlerModels.Response{
		Id:          id,
		Link:        link,
		Title:       title,
		Description: description,
		CreatedDate: createdDate,
		RootDomain:  rootDomain,
	}
}

func GetDbResponse() dbModels.Response {
	return dbModels.Response{
		Id:          id,
		Link:        link,
		Title:       title,
		Description: description,
		CreatedDate: createdDate,
		RootDomain:  rootDomain,
	}
}
