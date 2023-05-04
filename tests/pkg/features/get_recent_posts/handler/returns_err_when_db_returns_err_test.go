package handler

import (
	"context"
	"github.com/stretchr/testify/assert"
	dbError "github.com/whkelvin/stamp/pkg/features/errors/db"
	handlerError "github.com/whkelvin/stamp/pkg/features/errors/handler"
	dbModels "github.com/whkelvin/stamp/pkg/features/get_recent_posts/db/models"
	. "github.com/whkelvin/stamp/pkg/features/get_recent_posts/handler"
	handlerModels "github.com/whkelvin/stamp/pkg/features/get_recent_posts/handler/models"
	"testing"
)

var errorStr string = "db error"

type DbThatReturnsErrMock struct{}

func (m *DbThatReturnsErrMock) GetRecentPosts(ctx context.Context, req dbModels.Request) (*dbModels.Response, error) {
	return nil, dbError.New(errorStr)
}

func Test_Handler_Should_Return_Handler_Error_When_Db_Returns_Error(t *testing.T) {
	var handler IGetRecentPostsHandler = &GetRecentPostsHandler{
		DbService: &DbThatReturnsErrMock{}}

	_, err := handler.GetRecentPosts(context.Background(), handlerModels.Request{})
	assert.Equal(t, err, handlerError.New(errorStr, false))
}
