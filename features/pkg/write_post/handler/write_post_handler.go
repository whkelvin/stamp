package handler

import (
	. "github.com/whkelvin/stamp/features/pkg/write_post/models"
)

type IWritePostHandler interface {
	WritePost(req Request) (string, error)
}

type WritePostHandler struct {
	//GetUserDbService db.IGetUserDbService
}

func (handler *WritePostHandler) WritePost(req Request) (string, error) {
	//user, err := handler.GetUserDbService.GetUser(req.Id)

	//if err != nil {
	//	return models.GetUserResponse{}, err
	//}

	//return models.GetUserResponse{
	//	Id:   int(user.UserId),
	//	Name: user.Username,
	//}, nil
	return "ok", nil
}
