package controllers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	. "github.com/whkelvin/stamp/pkg/api/generated"
	handlerError "github.com/whkelvin/stamp/pkg/features/errors/handler"
	handlerModel "github.com/whkelvin/stamp/pkg/features/write_post/handler/models"
)

func parseWritePostRequest(c echo.Context) (*PostPostRequest, error) {
	var req PostPostRequest
	err := c.Bind(&req)
	if err != nil {
		return nil, err
	}

	if req.Link == "" {
		return nil, errors.New("field 'link' is required.")
	}

	if req.Title == "" {
		return nil, errors.New("field 'title' is required.")
	}

	if req.RootDomain == "" {
		return nil, errors.New("field 'rootDomain' is required.")
	}

	return &req, nil
}

func (server *ApiServer) CreatePost(ctx echo.Context) error {
	req, err := parseWritePostRequest(ctx)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	handlerReq := handlerModel.Request{
		Link:        req.Link,
		Title:       req.Title,
		Description: req.Description,
		RootDomain:  req.RootDomain,
	}
	dto, err := server.WritePostHandler.WritePost(ctx.Request().Context(), handlerReq)
	if err != nil {
		handlerErr, ok := err.(handlerError.HandlerError)
		if ok && handlerErr.IsBadInput() {
			return ctx.String(http.StatusBadRequest, err.Error())
		}
		return ctx.String(http.StatusInternalServerError, "Something went wrong, try again later.")
	}

	res := Post{
		Id:          dto.Id,
		Title:       dto.Title,
		Description: dto.Description,
		CreatedDate: dto.CreatedDate,
		Link:        dto.Link,
		RootDomain:  dto.RootDomain,
	}

	return ctx.JSON(http.StatusCreated, res)
}
