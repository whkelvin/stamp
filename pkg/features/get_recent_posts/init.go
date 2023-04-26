package get_recent_posts

import (
	. "github.com/whkelvin/stamp/pkg/features/get_recent_posts/db"
	. "github.com/whkelvin/stamp/pkg/features/get_recent_posts/handler"
	"go.mongodb.org/mongo-driver/mongo"
)

type GetRecentPostsFeature struct {
	MongoDbClient         *mongo.Client
	MongoDbDatabaseName   string
	MongoDbCollectionName string
}

func (feat *GetRecentPostsFeature) Init() *GetRecentPostsHandler {
	var getRecentPostsDbService *GetRecentPostsDbService = &GetRecentPostsDbService{MongoDbClient: feat.MongoDbClient, MongoDbDatabaseName: feat.MongoDbDatabaseName, MongoDbCollectionName: feat.MongoDbCollectionName}
	var getRecentPostsHandler *GetRecentPostsHandler = &GetRecentPostsHandler{DbService: getRecentPostsDbService}
	return getRecentPostsHandler
}
