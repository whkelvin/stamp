package db

import (
	"context"
	dbError "github.com/whkelvin/stamp/pkg/features/errors/db"
	"github.com/whkelvin/stamp/pkg/features/log_in/db/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type ILogInDbService interface {
	CreateOrGetUser(ctx context.Context, newPost models.Request) (*models.Response, error)
}

type LogInDbService struct {
	MongoDbClient         *mongo.Client
	MongoDbDatabaseName   string
	MongoDbCollectionName string
}

func (db *LogInDbService) CreateOrGetUser(ctx context.Context, request models.Request) (*models.Response, error) {
	var getUserReq getUserRequest = getUserRequest{
		AuthProviderName: request.AuthProviderName,
		Username:         request.Username,
	}

	user, err := db.getUser(ctx, getUserReq)
	if err != nil {
		return nil, err
	}

	if user != nil {
		var res models.Response = models.Response{
			IsAdmin:          user.IsAdmin,
			AuthProviderName: user.AuthProviderName,
			Username:         user.Username,
		}
		return &res, nil
	}

	var createUserReq createUserRequest = createUserRequest{
		IsAdmin:          false,
		Username:         request.Username,
		AuthProviderName: request.AuthProviderName,
	}
	result, err := db.createUser(ctx, createUserReq)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, dbError.New("user creation failed.")
	}

	var res models.Response = models.Response{
		IsAdmin:          result.IsAdmin,
		AuthProviderName: result.AuthProviderName,
		Username:         result.Username,
	}
	return &res, nil
}
