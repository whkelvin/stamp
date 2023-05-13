package refresh_token

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	. "github.com/whkelvin/stamp/pkg/api/controllers"
	handlerModels "github.com/whkelvin/stamp/pkg/features/refresh_token/handler/models"
)

type HandlerThatReturnsToken struct {
}

func (m *HandlerThatReturnsToken) RefreshToken(ctx context.Context, req handlerModels.Request) (*handlerModels.Response, error) {
	return &handlerModels.Response{
		Jwt: "helloworld",
	}, nil
}

func Test_Controller_Should_Return_Ok_When_Handler_Returns_Token(t *testing.T) {
	var e *echo.Echo = echo.New()

	rec := httptest.NewRecorder()
	token := []string{"Bearer 1234aosefj"}

	for i := 0; i < len(token); i++ {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
		req.Header.Add("Authorization", token[i])
		ctx := e.NewContext(req, rec)

		server := ApiServer{
			RefreshTokenHandler: &HandlerThatReturnsToken{},
		}

		server.RefreshToken(ctx)
		assert.Equal(t, rec.Code, http.StatusOK)
	}
}
