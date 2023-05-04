package write_post

import (
	"context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	. "github.com/whkelvin/stamp/pkg/api/controllers"
	. "github.com/whkelvin/stamp/pkg/api/generated"
	handlerModels "github.com/whkelvin/stamp/pkg/features/write_post/handler/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type HandlerThatReturnsResult struct {
}

func (m *HandlerThatReturnsResult) WritePost(ctx context.Context, req handlerModels.Request) (*handlerModels.Response, error) {
	return &handlerModels.Response{
		Id:          "6451167307ff6dafe2ea09ea",
		Title:       "test",
		Description: "test",
		Link:        "test",
		CreatedDate: "test",
		RootDomain:  "test",
	}, nil
}

func Test_Controller_Should_Return_Created_When_Handler_Returns_Result(t *testing.T) {
	var e *echo.Echo = echo.New()

	var reqBody PostPostRequest = PostPostRequest{
		Description: "test",
		Link:        "test",
		RootDomain:  "test",
		Title:       "test",
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
		WritePostHandler: &HandlerThatReturnsResult{},
	}

	server.CreatePost(ctx)

	assert.Equal(t, rec.Code, http.StatusCreated)
}
