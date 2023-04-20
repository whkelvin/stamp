package controller

import (
	"errors"
	"github.com/labstack/echo/v4"
	. "github.com/whkelvin/stamp/pkg/api/features/write_post/models"
	"github.com/whkelvin/stamp/pkg/features/write_post/handler"
	handlerModel "github.com/whkelvin/stamp/pkg/features/write_post/handler/models"
	"net/http"
)

type WritePostController struct {
	Handler handler.IWritePostHandler
}

func (c *WritePostController) Init(route string, e *echo.Echo) {
	e.POST(route, c.WritePost)
	e.POST("/api/v1/testpost", c.TestPost)
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

func (controller *WritePostController) TestPost(c echo.Context) error {
	req, err := parseWritePostRequest(c)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	handlerReq := handlerModel.Request{
		Link:        req.Link,
		Title:       req.Title,
		Description: req.Description,
	}

	return c.JSON(http.StatusCreated, handlerReq)
}
