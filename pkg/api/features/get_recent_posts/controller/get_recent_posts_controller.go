package controller

import (
	"errors"
	"github.com/labstack/echo/v4"
	. "github.com/whkelvin/stamp/pkg/api/features/get_recent_posts/models"
	"github.com/whkelvin/stamp/pkg/features/get_recent_posts/handler"
	handlerModel "github.com/whkelvin/stamp/pkg/features/get_recent_posts/handler/models"
	"net/http"
)

type GetRecentPostsController struct {
	Handler handler.IGetRecentPostsHandler
}

func (c *GetRecentPostsController) Init(route string, e *echo.Echo) {
	e.GET(route, c.GetRecentPosts)
}

func parseGetRecentPostsRequest(c echo.Context) (*Request, error) {

	var req Request
	err := c.Bind(&req)
	if err != nil {
		return nil, err
	}

	if req.Page < 1 {
		return nil, errors.New("field 'page' must be greater than 0.")
	}

	return &req, nil
}

func (controller *GetRecentPostsController) GetRecentPosts(c echo.Context) error {

	req, err := parseGetRecentPostsRequest(c)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	handlerReq := handlerModel.Request{
		Skip: (req.Page - 1) * req.Size,
		Take: req.Size,
	}
	result, err := controller.Handler.GetRecentPosts(c.Request().Context(), handlerReq)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Something went wrong, try again later.")
	}

	var posts []Post

	for i := 0; i < len(result.Posts); i++ {
		posts = append(posts, Post{
			PostId:      result.Posts[i].PostId,
			CreatedDate: result.Posts[i].CreatedDate,
			Title:       result.Posts[i].Title,
			Link:        result.Posts[i].Link,
			Description: result.Posts[i].Description,
			RootDomain:  result.Posts[i].RootDomain,
		})
	}

	res := Response{
		Count:    result.Count,
		PageSize: req.Size,
		Posts:    posts,
		Page:     req.Page,
	}

	return c.JSON(http.StatusCreated, res)
}
