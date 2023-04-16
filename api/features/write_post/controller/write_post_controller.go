package controller

import (
	//"fmt"
	"errors"
	"net/http"
	//"stamp/features/pkg/get_user/handler"
	//"stamp/features/pkg/get_user/models"
	//"strconv"

	"github.com/labstack/echo/v4"
	//"github.com/labstack/gommon/log"
	. "github.com/whkelvin/stamp/api/features/write_post/models"

	"github.com/whkelvin/stamp/features/pkg/write_post/handler"
	handlerModel "github.com/whkelvin/stamp/features/pkg/write_post/handler/models"
)

type WritePostController struct {
	Handler handler.IWritePostHandler
}

func (c *WritePostController) Init(route string, e *echo.Echo) {
	e.POST(route, c.WritePost)
}

func parseWritePostRequest(c echo.Context) (*Request, error) {

	var req Request
	err := c.Bind(&req)
	if err != nil {
		return nil, err
	}

	if req.Link == "" {
		return nil, errors.New("field 'link' is required")
	}

	if req.Title == "" {
		return nil, errors.New("field 'title' is required ")
	}

	return &req, nil
}

func (controller *WritePostController) WritePost(c echo.Context) error {

	req, err := parseWritePostRequest(c)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	handlerReq := handlerModel.Request{
		Link:        req.Link,
		Title:       req.Title,
		Description: req.Description,
	}
	err = controller.Handler.WritePost(handlerReq)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Something went wrong, try again later.")
	}

	return c.String(http.StatusCreated, "Post Created.")
}
