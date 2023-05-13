package refresh_token

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	. "github.com/whkelvin/stamp/pkg/api/controllers"
	handlerModels "github.com/whkelvin/stamp/pkg/features/refresh_token/handler/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type HandlerMock struct {
}

func (m *HandlerMock) RefreshToken(ctx context.Context, req handlerModels.Request) (*handlerModels.Response, error) {
	return nil, nil
}

func Test_Controller_Should_Return_Bad_Request_When_Token_Is_Invalid(t *testing.T) {
	var e *echo.Echo = echo.New()

	rec := httptest.NewRecorder()
	token := []string{"earer test", "random string here"}

	for i := 0; i < len(token); i++ {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
		req.Header.Add("Authorization", token[i])
		ctx := e.NewContext(req, rec)

		server := ApiServer{
			RefreshTokenHandler: &HandlerMock{},
		}

		server.RefreshToken(ctx)
		assert.Equal(t, rec.Code, http.StatusBadRequest)
	}
}
