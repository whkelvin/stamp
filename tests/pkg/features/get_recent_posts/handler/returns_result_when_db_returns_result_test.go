package handler

import (
	"context"
	"github.com/stretchr/testify/assert"
	dbModels "github.com/whkelvin/stamp/pkg/features/get_recent_posts/db/models"
	. "github.com/whkelvin/stamp/pkg/features/get_recent_posts/handler"
	handlerModels "github.com/whkelvin/stamp/pkg/features/get_recent_posts/handler/models"
	"testing"
)

type HappyPathDbMock struct{}

func (m *HappyPathDbMock) GetRecentPosts(ctx context.Context, req dbModels.Request) (*dbModels.Response, error) {
	res := getDbModelResponse()
	return &res, nil
}

func Test_Handler_Should_Return_Result_When_Db_Returns_Result(t *testing.T) {
	var handler IGetRecentPostsHandler = &GetRecentPostsHandler{
		DbService: &HappyPathDbMock{}}

	actual, err := handler.GetRecentPosts(context.Background(), handlerModels.Request{})
	assert.Equal(t, err, nil)

	expected := handlerModels.Response{
		Count: getDbModelResponse().Count,
		Posts: []handlerModels.Post{
			getHandlerModelPost(),
		},
	}

	assert.Equal(t, expected, *actual)
}
