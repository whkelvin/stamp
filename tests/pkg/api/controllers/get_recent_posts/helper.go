package get_recent_posts

import (
	"encoding/json"
	"time"

	controllerModels "github.com/whkelvin/stamp/pkg/api/generated"
	handlerModels "github.com/whkelvin/stamp/pkg/features/get_recent_posts/handler/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var rawId string = "6451167307ff6dafe2ea09ea"
var id, _ = primitive.ObjectIDFromHex(rawId)
var link = "test link"
var title = "Test Title"
var description = "Test Description"
var createdDate = time.Now().UTC()
var rootDomain = "Test Domain"

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

func getControllerModelPost() controllerModels.Post {
	return controllerModels.Post{
		CreatedDate: createdDate.Format(time.RFC3339),
		RootDomain:  rootDomain,
		Description: description,
		Id:          rawId,
		Link:        link,
		Title:       title,
	}
}

func getControllerResponseInJson() string {
	res := controllerModels.PostResultSet{
		Count:    1,
		PageSize: 1,
		Posts: []controllerModels.Post{
			getControllerModelPost(),
		},
	}
	json, _ := json.Marshal(res)
	return string(json)
}
