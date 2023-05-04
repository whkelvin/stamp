package get_recent_posts

import (
	"context"
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

type HandlerThatReturnsResult struct {
}

func (m *HandlerThatReturnsResult) GetRecentPosts(ctx context.Context, req handlerModels.Request) (*handlerModels.Response, error) {
	return &handlerModels.Response{
		Count: 1,
		Posts: []handlerModels.Post{
			getHandlerModelPost(),
		},
	}, nil
}

func Test_Controller_Should_Return_Result_When_Handler_Returns_Result(t *testing.T) {
	var one int32 = 1
	var mtString = ""
	var e *echo.Echo = echo.New()

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	ctx := e.NewContext(req, rec)

	server := ApiServer{
		GetRecentPostsHandler: &HandlerThatReturnsResult{},
	}

	server.GetRecentPosts(ctx, GetRecentPostsParams{
		LastFetchedItemId: &mtString,
		Size:              &one,
	})

	assert.Equal(t, rec.Code, http.StatusOK)
	assert.Equal(t, strings.TrimSuffix(rec.Body.String(), "\n"), getControllerResponseInJson())
}
