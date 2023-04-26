package db

import (
	"context"
	. "github.com/whkelvin/stamp/pkg/features/get_recent_posts/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IGetRecentPostsDbService interface {
	GetRecentPosts(ctx context.Context, req Request) (*Response, error)
}

type GetRecentPostsDbService struct {
	MongoDbClient         *mongo.Client
	MongoDbDatabaseName   string
	MongoDbCollectionName string
}

func (db *GetRecentPostsDbService) GetRecentPosts(ctx context.Context, req Request) (*Response, error) {
	coll := db.MongoDbClient.Database(db.MongoDbDatabaseName).Collection(db.MongoDbCollectionName)

	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{primitive.E{Key: "created_date", Value: -1}}).SetLimit(int64(req.Take)).SetSkip(int64(req.Skip))

	cursor, err := coll.Find(ctx, filter, opts)

	var result []Post
	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	res := &Response{
		Posts: result,
		Count: len(result),
	}
	return res, nil
}
