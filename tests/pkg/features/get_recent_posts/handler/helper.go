package handler

import (
	dbModels "github.com/whkelvin/stamp/pkg/features/get_recent_posts/db/models"
	handlerModels "github.com/whkelvin/stamp/pkg/features/get_recent_posts/handler/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var rawId string = "6451167307ff6dafe2ea09ea"
var id, _ = primitive.ObjectIDFromHex(rawId)
var link = "test link"
var title = "Test Title"
var description = "Test Description"
var createdDate = time.Now().UTC()
var rootDomain = "Test Domain"

func getDbModelPost() dbModels.Post {
	return dbModels.Post{
		Id:          id,
		Link:        link,
		Title:       title,
		Description: description,
		CreatedDate: createdDate,
		RootDomain:  rootDomain,
	}
}

func getDbModelResponse() dbModels.Response {
	return dbModels.Response{
		Count: 1,
		Posts: []dbModels.Post{
			getDbModelPost(),
		},
	}
}

func getHandlerModelPost() handlerModels.Post {
	return handlerModels.Post{
		Id:          rawId,
		Link:        link,
		Title:       title,
		Description: description,
		CreatedDate: createdDate,
		RootDomain:  rootDomain,
	}
}
