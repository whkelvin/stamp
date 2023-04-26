package write_post

import (
	. "github.com/whkelvin/stamp/pkg/features/write_post/db"
	. "github.com/whkelvin/stamp/pkg/features/write_post/handler"
	"go.mongodb.org/mongo-driver/mongo"
)

type WritePostFeature struct {
	MongoDbClient         *mongo.Client
	MongoDbDatabaseName   string
	MongoDbCollectionName string
}

func (feat *WritePostFeature) Init() *WritePostHandler {
	var writePostDbService *WritePostDbService = &WritePostDbService{MongoDbClient: feat.MongoDbClient, MongoDbDatabaseName: feat.MongoDbDatabaseName, MongoDbCollectionName: feat.MongoDbCollectionName}
	var writePostHandler *WritePostHandler = &WritePostHandler{DbService: writePostDbService}
	return writePostHandler
}
