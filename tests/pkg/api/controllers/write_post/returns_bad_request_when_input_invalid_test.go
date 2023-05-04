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

type WritePostHandlerMock struct {
}

func (m *WritePostHandlerMock) WritePost(ctx context.Context, req handlerModels.Request) (*handlerModels.Response, error) {
	return &handlerModels.Response{}, nil
}

func Test_Controller_Should_Return_Bad_Request_When_Input_Invalid(t *testing.T) {
	testCases := []TestCase{
		{Link: "", Title: "not empty", RootDomain: "not empty", Description: "not empty", Expected: http.StatusBadRequest},
		{Link: "not empty", Title: "", RootDomain: "not empty", Description: "not empty", Expected: http.StatusBadRequest},
		{Link: "not empty", Title: "not empty", RootDomain: "", Description: "not empty", Expected: http.StatusBadRequest},
	}

	for i := 0; i < len(testCases); i++ {
		var e *echo.Echo = echo.New()

		var reqBody PostPostRequest = PostPostRequest{
			Description: testCases[i].Description,
			Link:        testCases[i].Link,
			RootDomain:  testCases[i].RootDomain,
			Title:       testCases[i].Title,
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
			WritePostHandler: &WritePostHandlerMock{},
		}

		server.CreatePost(ctx)

		assert.Equal(t, rec.Code, testCases[i].Expected)
	}
}
