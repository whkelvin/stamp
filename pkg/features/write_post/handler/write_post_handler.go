package handler

import (
	"context"
	"errors"
	. "github.com/whkelvin/stamp/pkg/features/write_post/db"
	dbModels "github.com/whkelvin/stamp/pkg/features/write_post/db/models"
	. "github.com/whkelvin/stamp/pkg/features/write_post/handler/models"
	"github.com/whkelvin/stamp/pkg/helpers"
	"time"
)

type IWritePostHandler interface {
	WritePost(ctx context.Context, req Request) (*Response, error)
}

type WritePostHandler struct {
	DbService IWritePostDbService
}

func (handler *WritePostHandler) WritePost(ctx context.Context, req Request) (*Response, error) {

	newPost := dbModels.Request{
		Link:        req.Link,
		Title:       req.Title,
		Description: req.Description,
		RootDomain:  req.RootDomain,
	}

	if newPost.RootDomain == "youtube.com" {
		result, err := helpers.GetYoutubeEmbedLink(newPost.Link)
		if err != nil {
			return nil, errors.New("Invalid youtube link")
		}
		newPost.Link = result
	}

	dto, err := handler.DbService.CreatePost(ctx, newPost)
	if err != nil {
		return nil, err
	}

	var res *Response = &Response{
		PostId:      dto.PostId,
		Title:       dto.Title,
		Description: dto.Description,
		Link:        dto.Link,
		RootDomain:  dto.RootDomain,
		CreatedDate: dto.CreatedDate.UTC().Format(time.RFC3339),
	}

	return res, nil
}
