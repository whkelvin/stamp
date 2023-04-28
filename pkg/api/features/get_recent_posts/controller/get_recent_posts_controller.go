package controller

import (
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	. "github.com/whkelvin/stamp/pkg/api/features/get_recent_posts/models"
	gen "github.com/whkelvin/stamp/pkg/api/generated/models"
	"github.com/whkelvin/stamp/pkg/features/get_recent_posts/handler"
	handlerModel "github.com/whkelvin/stamp/pkg/features/get_recent_posts/handler/models"
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

	if req.Size == nil {
		var defaultSize int32 = 5
		req.Size = &defaultSize
	}

	if *req.Size < 1 {
		return nil, errors.New("field 'size' must be greater than or equal to 1.")
	}

	if req.Page == nil {
		var defaultPage int32 = 1
		req.Page = &defaultPage
	}

	if *req.Page < 1 {
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
		Skip: int((*req.Page - 1) * (*req.Size)),
		Take: int(*req.Size),
	}
	result, err := controller.Handler.GetRecentPosts(c.Request().Context(), handlerReq)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Something went wrong, try again later.")
	}

	var posts []gen.Post = []gen.Post{}

	for i := 0; i < len(result.Posts); i++ {
		posts = append(posts, gen.Post{
			Id:          result.Posts[i].Id,
			CreatedDate: result.Posts[i].CreatedDate.UTC().Format(time.RFC3339),
			Title:       result.Posts[i].Title,
			Link:        result.Posts[i].Link,
			Description: result.Posts[i].Description,
			RootDomain:  result.Posts[i].RootDomain,
		})
	}

	res := Response{
		Count:    int32(len(posts)),
		PageSize: *req.Size,
		Posts:    posts,
		Page:     *req.Page,
	}

	return c.JSON(http.StatusOK, res)
}
