package controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	. "github.com/whkelvin/stamp/pkg/api/generated"
	handlerModel "github.com/whkelvin/stamp/pkg/features/get_recent_posts/handler/models"
)

func parseParams(params GetRecentPostsParams) (string, int32, error) {
	var lastFetchedItemId string = ""
	var size int32 = 10

	if params.LastFetchedItemId != nil {
		lastFetchedItemId = *params.LastFetchedItemId
	}

	if params.Size != nil {
		size = *params.Size
		if size <= 0 {
			return "", 0, errors.New("Field 'size' has to be greater than zero.")
		}
	}
	return lastFetchedItemId, size, nil
}

func (server *ApiServer) GetRecentPosts(ctx echo.Context, params GetRecentPostsParams) error {

	lastFetchedItemId, size, err := parseParams(params)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	handlerReq := handlerModel.Request{
		LastFetchedItemId: lastFetchedItemId,
		Take:              int(size),
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
	}

	return ctx.JSON(http.StatusOK, res)
}
