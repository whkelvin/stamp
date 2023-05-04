package handler

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	handlerError "github.com/whkelvin/stamp/pkg/features/errors/handler"
	dbModels "github.com/whkelvin/stamp/pkg/features/write_post/db/models"
	. "github.com/whkelvin/stamp/pkg/features/write_post/handler"
	handlerModels "github.com/whkelvin/stamp/pkg/features/write_post/handler/models"
)

type DbThatReturnsError struct{}

func (m *DbThatReturnsError) CreatePost(ctx context.Context, req dbModels.Request) (*dbModels.Response, error) {
	return nil, errors.New("")
}

func Test_Handler_Should_Return_Generic_Error_When_Db_Returns_Error(t *testing.T) {
	var handler IWritePostHandler = &WritePostHandler{
		DbService: &DbThatReturnsError{}}

	_, err := handler.WritePost(context.Background(), handlerModels.Request{
		Link:        "https://youtu.be/c4OyfL5o7DU",
		RootDomain:  "youtube.com",
		Title:       "irrelevent here",
		Description: "irrelevent here",
	})

	handlerErr, ok := err.(handlerError.HandlerError)
	if !ok || handlerErr.IsBadInput() {
		assert.Fail(t, "handler did not return Generic Handler Error")
	}
}
