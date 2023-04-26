package handler

import (
	"context"

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
		Take: req.Take,
		Skip: req.Skip,
	})
	if err != nil {
		return nil, err
	}

	var posts []handlerModels.Post
	for i := 0; i < len(dto.Posts); i++ {
		dest := handlerModels.Post{
			PostId:      dto.Posts[i].PostId.Hex(),
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
