package handler

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	dbModels "github.com/whkelvin/stamp/pkg/features/write_post/db/models"
	. "github.com/whkelvin/stamp/pkg/features/write_post/handler"
)

type DbThatReturnsResult struct{}

func (m *DbThatReturnsResult) CreatePost(ctx context.Context, req dbModels.Request) (*dbModels.Response, error) {
	res := GetDbResponse()
	return &res, nil
}

func Test_Handler_Should_Return_Result_When_Db_Returns_Result(t *testing.T) {
	var handler IWritePostHandler = &WritePostHandler{
		DbService: &DbThatReturnsResult{}}

	res, err := handler.WritePost(context.Background(), GetHandlerRequest())

	assert.Equal(t, err, nil)
	assert.Equal(t, *res, GetHandlerResponse())
}
