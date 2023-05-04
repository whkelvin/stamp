package get_recent_posts

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	. "github.com/whkelvin/stamp/pkg/api/controllers"
	. "github.com/whkelvin/stamp/pkg/api/generated"
	handlerModels "github.com/whkelvin/stamp/pkg/features/get_recent_posts/handler/models"
)

type HandlerThatReturnsNil struct {
}

func (m *HandlerThatReturnsNil) GetRecentPosts(ctx context.Context, req handlerModels.Request) (*handlerModels.Response, error) {
	return nil, nil
}

func Test_Controller_Should_Return_Empty_Result_When_Handler_Returns_Null(t *testing.T) {
	var one int32 = 1
	var mtString = ""
	var e *echo.Echo = echo.New()

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	ctx := e.NewContext(req, rec)

	server := ApiServer{
		GetRecentPostsHandler: &HandlerThatReturnsNil{},
	}

	server.GetRecentPosts(ctx, GetRecentPostsParams{
		LastFetchedItemId: &mtString,
		Size:              &one,
	})

	res := PostResultSet{
		Count:    0,
		PageSize: one,
		Posts:    []Post{},
	}
	json, err := json.Marshal(res)
	if err != nil {
		assert.Fail(t, "failed to convert object to json")
	}

	assert.Equal(t, rec.Code, http.StatusOK)

	assert.Equal(t, strings.TrimSuffix(rec.Body.String(), "\n"), string(json))
}
