package handler

import (
	"context"
	"github.com/stretchr/testify/assert"
	dbModels "github.com/whkelvin/stamp/pkg/features/get_recent_posts/db/models"
	. "github.com/whkelvin/stamp/pkg/features/get_recent_posts/handler"
	handlerModels "github.com/whkelvin/stamp/pkg/features/get_recent_posts/handler/models"
	"testing"
)

type DbThatAlwaysReturnNil struct{}

func (m *DbThatAlwaysReturnNil) GetRecentPosts(ctx context.Context, req dbModels.Request) (*dbModels.Response, error) {
	return nil, nil
}

func Test_Handler_Should_Return_Empty_Result_When_Db_Returns_Null(t *testing.T) {
	var handler IGetRecentPostsHandler = &GetRecentPostsHandler{
		DbService: &DbThatAlwaysReturnNil{}}

	actual, err := handler.GetRecentPosts(context.Background(), handlerModels.Request{})
	assert.Equal(t, err, nil)

	expected := handlerModels.Response{
		Count: 0,
		Posts: []handlerModels.Post{},
	}

	assert.Equal(t, expected, *actual)
}
