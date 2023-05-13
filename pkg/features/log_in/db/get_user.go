package db

import (
	"context"
	dbError "github.com/whkelvin/stamp/pkg/features/errors/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type getUserRequest struct {
	AuthProviderName string
	Username         string
}

type getUserResponse struct {
	Id               string
	IsAdmin          bool
	AuthProviderName string
	Username         string
}

func (db *LogInDbService) getUser(ctx context.Context, request getUserRequest) (*getUserResponse, error) {
	coll := db.MongoDbClient.Database(db.MongoDbDatabaseName).Collection(db.MongoDbCollectionName)

	filter := bson.D{
		primitive.E{
			Key: "authProviders",
			Value: bson.D{
				primitive.E{Key: "name", Value: request.AuthProviderName},
				primitive.E{Key: "username", Value: request.Username},
			},
		},
	}
	opts := options.FindOne()
	res := coll.FindOne(ctx, filter, opts)

	var result userSchema
	err := res.Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, dbError.New("cannot decode user object")
	}

	var response getUserResponse = getUserResponse{
		Id:               result.Id.Hex(),
		IsAdmin:          result.IsAdmin,
		Username:         request.Username,
		AuthProviderName: request.AuthProviderName,
	}
	return &response, nil
}
