package write_post

import (
	"context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	. "github.com/whkelvin/stamp/pkg/api/controllers"
	. "github.com/whkelvin/stamp/pkg/api/generated"
	handlerError "github.com/whkelvin/stamp/pkg/features/errors/handler"
	handlerModels "github.com/whkelvin/stamp/pkg/features/write_post/handler/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type HandlerThatReturnsBadInputError struct {
}

func (m *HandlerThatReturnsBadInputError) WritePost(ctx context.Context, req handlerModels.Request) (*handlerModels.Response, error) {
	return nil, handlerError.New("", true)
}

func Test_Controller_Should_Return_Bad_Request_When_Handler_Returns_Bad_Input_Error(t *testing.T) {
	var e *echo.Echo = echo.New()

	var reqBody PostPostRequest = PostPostRequest{
		Description: "",
		Link:        "not empty",
		RootDomain:  "not empty",
		Title:       "not empty",
	}
	json, err := json.Marshal(reqBody)
	if err != nil {
		assert.Fail(t, "json generation failed")
	}

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(json)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx := e.NewContext(req, rec)

	server := ApiServer{
		WritePostHandler: &HandlerThatReturnsBadInputError{},
	}

	server.CreatePost(ctx)

	assert.Equal(t, rec.Code, http.StatusBadRequest)
}
