package controllers_test

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	. "github.com/whkelvin/stamp/pkg/api/controllers"
	. "github.com/whkelvin/stamp/pkg/api/generated"
	handlerModels "github.com/whkelvin/stamp/pkg/features/get_recent_posts/handler/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

type GetRecentPostsHandlerMock struct {
}

func (m *GetRecentPostsHandlerMock) GetRecentPosts(ctx context.Context, req handlerModels.Request) (*handlerModels.Response, error) {
	return &handlerModels.Response{}, nil
}

type TestCase struct {
	LastFetchedId *string
	Size          *int32
	Expected      int
}

func TestControllerInputValidation(t *testing.T) {
	var zero int32 = 0
	var emptyString = ""

	testCases := []TestCase{
		{LastFetchedId: &emptyString, Size: &zero, Expected: http.StatusBadRequest},
	}

	for i := 0; i < len(testCases); i++ {
		var e *echo.Echo = echo.New()

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		ctx := e.NewContext(req, rec)

		server := ApiServer{
			GetRecentPostsHandler: &GetRecentPostsHandlerMock{},
		}

		server.GetRecentPosts(ctx, GetRecentPostsParams{
			LastFetchedItemId: testCases[i].LastFetchedId,
			Size:              testCases[i].Size,
		})

		assert.Equal(t, rec.Code, testCases[i].Expected)
	}
}

type GetRecentPostsHandlerNilMock struct {
}

func (m *GetRecentPostsHandlerNilMock) GetRecentPosts(ctx context.Context, req handlerModels.Request) (*handlerModels.Response, error) {
	return nil, nil
}

func TestControllerShouldReturnWhenGetRecentPostsReturnsNull(t *testing.T) {
	var one int32 = 1
	var mtString = ""
	var e *echo.Echo = echo.New()

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	ctx := e.NewContext(req, rec)

	server := ApiServer{
		GetRecentPostsHandler: &GetRecentPostsHandlerNilMock{},
	}

	server.GetRecentPosts(ctx, GetRecentPostsParams{
		LastFetchedItemId: &mtString,
		Size:              &one,
	})

	assert.Equal(t, rec.Code, http.StatusOK)
}