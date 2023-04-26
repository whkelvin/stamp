package db

import (
	"context"
	"github.com/whkelvin/stamp/pkg/features/write_post/db/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type IWritePostDbService interface {
	CreatePost(ctx context.Context, newPost models.Request) (*models.Response, error)
}

type WritePostDbService struct {
	MongoDbClient         *mongo.Client
	MongoDbDatabaseName   string
	MongoDbCollectionName string
}

func (db *WritePostDbService) CreatePost(ctx context.Context, request models.Request) (*models.Response, error) {
	coll := db.MongoDbClient.Database(db.MongoDbDatabaseName).Collection(db.MongoDbCollectionName)
	request.CreatedDate = time.Now().UTC().Format(time.RFC3339)

	opts := options.InsertOne().SetBypassDocumentValidation(true)

	result, err := coll.InsertOne(ctx, request, opts)
	if err != nil {
		return nil, err
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	var response models.Response = models.Response{
		PostId:      id,
		CreatedDate: request.CreatedDate,
		Link:        request.Link,
		Description: request.Description,
		Title:       request.Description,
		RootDomain:  request.RootDomain,
	}

	return &response, nil
}
