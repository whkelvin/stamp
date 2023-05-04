package get_recent_posts

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	. "github.com/whkelvin/stamp/pkg/api/controllers"
	. "github.com/whkelvin/stamp/pkg/api/generated"
	handlerError "github.com/whkelvin/stamp/pkg/features/errors/handler"
	handlerModels "github.com/whkelvin/stamp/pkg/features/get_recent_posts/handler/models"
)

type HandlerThatReturnsError struct {
}

func (m *HandlerThatReturnsError) GetRecentPosts(ctx context.Context, req handlerModels.Request) (*handlerModels.Response, error) {
	return nil, handlerError.New("", false)
}

func Test_Controller_Should_Return_Internal_Server_Error_When_Handler_Returns_Error(t *testing.T) {
	var e *echo.Echo = echo.New()

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	ctx := e.NewContext(req, rec)

	server := ApiServer{
		GetRecentPostsHandler: &HandlerThatReturnsError{},
	}

	var id string = ""
	var size int32 = 1
	server.GetRecentPosts(ctx, GetRecentPostsParams{
		LastFetchedItemId: &id,
		Size:              &size,
	})

	assert.Equal(t, rec.Code, http.StatusInternalServerError)
}
