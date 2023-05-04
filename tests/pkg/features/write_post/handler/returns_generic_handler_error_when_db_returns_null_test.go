package handler

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	handlerError "github.com/whkelvin/stamp/pkg/features/errors/handler"
	dbModels "github.com/whkelvin/stamp/pkg/features/write_post/db/models"
	. "github.com/whkelvin/stamp/pkg/features/write_post/handler"
	handlerModels "github.com/whkelvin/stamp/pkg/features/write_post/handler/models"
)

type DbThatReturnsNil struct{}

func (m *DbThatReturnsNil) CreatePost(ctx context.Context, req dbModels.Request) (*dbModels.Response, error) {
	return nil, nil
}

func Test_Handler_Should_Return_Generic_Handler_Error_When_Db_Returns_Nil(t *testing.T) {
	var handler IWritePostHandler = &WritePostHandler{
		DbService: &DbThatReturnsNil{}}

	_, err := handler.WritePost(context.Background(), handlerModels.Request{
		Link:        "https://youtu.be/c4OyfL5o7DU",
		RootDomain:  "youtube.com",
		Title:       "irrelevent here",
		Description: "irrelevent here",
	})

	handlerErr, ok := err.(handlerError.HandlerError)
	assert.True(t, ok)
	assert.False(t, handlerErr.IsBadInput())
}
