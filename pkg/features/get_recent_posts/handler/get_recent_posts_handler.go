package handler

import (
	"context"

	handlerError "github.com/whkelvin/stamp/pkg/features/errors/handler"
	. "github.com/whkelvin/stamp/pkg/features/get_recent_posts/db"
	dbModels "github.com/whkelvin/stamp/pkg/features/get_recent_posts/db/models"
	handlerModels "github.com/whkelvin/stamp/pkg/features/get_recent_posts/handler/models"
)

type IGetRecentPostsHandler interface {
	GetRecentPosts(ctx context.Context, req handlerModels.Request) (*handlerModels.Response, error)
}

type GetRecentPostsHandler struct {
	DbService IGetRecentPostsDbService
}

func (handler *GetRecentPostsHandler) GetRecentPosts(ctx context.Context, req handlerModels.Request) (*handlerModels.Response, error) {

	dto, err := handler.DbService.GetRecentPosts(ctx, dbModels.Request{
		Take:              req.Take,
		LastFetchedItemId: req.LastFetchedItemId,
	})
	if err != nil {
		return nil, handlerError.New(err.Error(), false)
	}

	var posts []handlerModels.Post = []handlerModels.Post{}

	if dto == nil {
		return &handlerModels.Response{
			Count: 0,
			Posts: posts,
		}, nil
	}

	for i := 0; i < dto.Count; i++ {
		dest := handlerModels.Post{
			Id:          dto.Posts[i].Id.Hex(),
			Description: dto.Posts[i].Description,
			Link:        dto.Posts[i].Link,
			Title:       dto.Posts[i].Title,
			CreatedDate: dto.Posts[i].CreatedDate,
			RootDomain:  dto.Posts[i].RootDomain,
		}
		posts = append(posts, dest)
	}

	var res *handlerModels.Response = &handlerModels.Response{
		Count: dto.Count,
		Posts: posts,
	}

	return res, nil
}
