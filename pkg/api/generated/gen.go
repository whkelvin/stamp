// Package models provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package generated

import (
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/labstack/echo/v4"
	"net/http"
)

const (
	Api_keyScopes = "api_key.Scopes"
)

// Post defines model for Post.
type Post struct {
	CreatedDate string `json:"createdDate"`
	Description string `json:"description"`
	Id          string `json:"id"`
	Link        string `json:"link"`
	RootDomain  string `json:"rootDomain"`
	Title       string `json:"title"`
}

// PostPostRequest defines model for PostPostRequest.
type PostPostRequest struct {
	Description string `json:"description"`
	Link        string `json:"link"`
	RootDomain  string `json:"rootDomain"`
	Title       string `json:"title"`
}

// PostResultSet defines model for PostResultSet.
type PostResultSet struct {
	Count    int32  `json:"count"`
	Page     int32  `json:"page"`
	PageSize int32  `json:"pageSize"`
	Posts    []Post `json:"posts"`
}

// GetRecentPostsParams defines parameters for GetRecentPosts.
type GetRecentPostsParams struct {
	// Size Number of results are included in a page
	Size *int32 `form:"size,omitempty" json:"size,omitempty"`

	// Page Page number
	Page *int32 `form:"page,omitempty" json:"page,omitempty"`
}

// CreatePostJSONRequestBody defines body for CreatePost for application/json ContentType.
type CreatePostJSONRequestBody = PostPostRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// create a new post
	// (POST /post)
	CreatePost(ctx echo.Context) error
	// get a list of most recent posts
	// (GET /posts)
	GetRecentPosts(ctx echo.Context, params GetRecentPostsParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// CreatePost converts echo context to params.
func (w *ServerInterfaceWrapper) CreatePost(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreatePost(ctx)
	return err
}

// GetRecentPosts converts echo context to params.
func (w *ServerInterfaceWrapper) GetRecentPosts(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetRecentPostsParams
	// ------------- Optional query parameter "size" -------------

	err = runtime.BindQueryParameter("form", true, false, "size", ctx.QueryParams(), &params.Size)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter size: %s", err))
	}

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetRecentPosts(ctx, params)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/post", wrapper.CreatePost)
	router.GET(baseURL+"/posts", wrapper.GetRecentPosts)

}
