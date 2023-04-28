package controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	. "github.com/whkelvin/stamp/pkg/api/generated"
	handlerModel "github.com/whkelvin/stamp/pkg/features/get_recent_posts/handler/models"
)

func parseParams(params GetRecentPostsParams) (int32, int32, error) {
	var page int32 = 1
	var size int32 = 10

	if params.Page != nil {
		page = *params.Page
		if page < 1 {
			return 0, 0, errors.New("Field 'page' has to be greater than or equal to 1.")
		}
	}

	if params.Size != nil {
		size = *params.Size
		if size <= 0 {
			return 0, 0, errors.New("Field 'size' has to be greater than zero.")
		}
	}
	return page, size, nil
}

func (server *ApiServer) GetRecentPosts(ctx echo.Context, params GetRecentPostsParams) error {

	page, size, err := parseParams(params)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	handlerReq := handlerModel.Request{
		Skip: int((page - 1) * (size)),
		Take: int(size),
	}

	result, err := server.GetRecentPostsHandler.GetRecentPosts(ctx.Request().Context(), handlerReq)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Something went wrong, try again later.")
	}

	if result == nil {
		res := PostResultSet{
			Count:    0,
			PageSize: size,
			Posts:    []Post{},
			Page:     page,
		}
		return ctx.JSON(http.StatusOK, res)
	}

	var posts []Post = []Post{}
	for i := 0; i < len(result.Posts); i++ {
		posts = append(posts, Post{
			Id:          result.Posts[i].Id,
			CreatedDate: result.Posts[i].CreatedDate.UTC().Format(time.RFC3339),
			Title:       result.Posts[i].Title,
			Link:        result.Posts[i].Link,
			Description: result.Posts[i].Description,
			RootDomain:  result.Posts[i].RootDomain,
		})
	}

	res := PostResultSet{
		Count:    int32(len(posts)),
		PageSize: size,
		Posts:    posts,
		Page:     page,
	}

	return ctx.JSON(http.StatusOK, res)
}
