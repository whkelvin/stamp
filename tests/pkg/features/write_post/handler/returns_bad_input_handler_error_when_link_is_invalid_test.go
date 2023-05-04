package handler

import (
	"context"
	"github.com/stretchr/testify/assert"
	handlerError "github.com/whkelvin/stamp/pkg/features/errors/handler"
	dbModels "github.com/whkelvin/stamp/pkg/features/write_post/db/models"
	. "github.com/whkelvin/stamp/pkg/features/write_post/handler"
	handlerModels "github.com/whkelvin/stamp/pkg/features/write_post/handler/models"
	"testing"
)

type DbMock struct{}

func (m *DbMock) CreatePost(ctx context.Context, req dbModels.Request) (*dbModels.Response, error) {
	return nil, nil
}

func Test_Handler_Should_Return_Bad_Input_Error_When_Link_Is_Invalid(t *testing.T) {
	var handler IWritePostHandler = &WritePostHandler{
		DbService: &DbMock{}}

	_, err := handler.WritePost(context.Background(), handlerModels.Request{
		Link:        "obviously not a youtube link",
		RootDomain:  "youtube.com",
		Title:       "irrelevent here",
		Description: "irrelevent here",
	})

	handlerErr, ok := err.(handlerError.HandlerError)
	if !ok || !handlerErr.IsBadInput() {
		assert.Fail(t, "handler did not return Bad Input Error")
	}
}
