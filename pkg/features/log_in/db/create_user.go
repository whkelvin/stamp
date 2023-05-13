package db

import (
	"context"
	dbError "github.com/whkelvin/stamp/pkg/features/errors/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type createUserRequest struct {
	IsAdmin          bool
	AuthProviderName string
	Username         string
}

type createUserResponse struct {
	Id               string
	IsAdmin          bool
	AuthProviderName string
	Username         string
}

func (db *LogInDbService) createUser(ctx context.Context, request createUserRequest) (*createUserResponse, error) {
	coll := db.MongoDbClient.Database(db.MongoDbDatabaseName).Collection(db.MongoDbCollectionName)

	newUser := createUserSchema{
		IsAdmin: false,
		AuthProviders: []authProvider{
			{
				Name:     request.AuthProviderName,
				Username: request.Username,
			},
		},
		CreatedDate: time.Now().UTC().Format(time.RFC3339),
	}

	opts := options.InsertOne().SetBypassDocumentValidation(true)

	result, err := coll.InsertOne(ctx, newUser, opts)
	if err != nil {
		return nil, dbError.New(err.Error())
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	var response createUserResponse = createUserResponse{
		Id:               id,
		IsAdmin:          false,
		Username:         request.Username,
		AuthProviderName: request.AuthProviderName,
	}
	return &response, nil
}
