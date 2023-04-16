package handler

import (
	. "github.com/whkelvin/stamp/features/pkg/write_post/db"
	dbModels "github.com/whkelvin/stamp/features/pkg/write_post/db/models"
	. "github.com/whkelvin/stamp/features/pkg/write_post/handler/models"
)

type IWritePostHandler interface {
	WritePost(req Request) error
}

type WritePostHandler struct {
	DbService IWritePostDbService
}

func (handler *WritePostHandler) WritePost(req Request) error {

	newPost := dbModels.NewPost{
		Link:        req.Link,
		Title:       req.Title,
		Description: req.Description,
	}

	handler.DbService.CreatePost(newPost)

	return nil
}
