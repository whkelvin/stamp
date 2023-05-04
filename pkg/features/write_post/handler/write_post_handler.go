package handler

import (
	"context"
	handlerError "github.com/whkelvin/stamp/pkg/features/errors/handler"
	. "github.com/whkelvin/stamp/pkg/features/write_post/db"
	dbModels "github.com/whkelvin/stamp/pkg/features/write_post/db/models"
	. "github.com/whkelvin/stamp/pkg/features/write_post/handler/models"
	"github.com/whkelvin/stamp/pkg/helpers"
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
			return nil, handlerError.New("Invalid youtube link", true)
		}
		newPost.Link = result
	}

	if newPost.RootDomain == "github.com" {
		err := helpers.ValidateGithubLink(newPost.Link)

		if err != nil {
			return nil, handlerError.New("Invalid github link", true)
		}
	}

	dto, err := handler.DbService.CreatePost(ctx, newPost)
	if err != nil {
		return nil, handlerError.New(err.Error(), false)
	}
	if dto == nil {
		return nil, handlerError.New("db failed to create post", false)
	}

	var res *Response = &Response{
		Id:          dto.Id,
		Title:       dto.Title,
		Description: dto.Description,
		Link:        dto.Link,
		RootDomain:  dto.RootDomain,
		CreatedDate: dto.CreatedDate,
	}

	return res, nil
}
