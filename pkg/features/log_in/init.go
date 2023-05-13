package log_in

import (
	. "github.com/whkelvin/stamp/pkg/features/log_in/db"
	. "github.com/whkelvin/stamp/pkg/features/log_in/handler"
	. "github.com/whkelvin/stamp/pkg/features/log_in/helpers"
	"go.mongodb.org/mongo-driver/mongo"
)

type LogInFeature struct {
	JwtSecret             string
	MongoDbClient         *mongo.Client
	MongoDbDatabaseName   string
	MongoDbCollectionName string
}

func (feat *LogInFeature) Init() *LogInHandler {
	var githubTokenValidator *GithubTokenValidator = &GithubTokenValidator{}

	var logInDbService *LogInDbService = &LogInDbService{
		MongoDbClient:         feat.MongoDbClient,
		MongoDbDatabaseName:   feat.MongoDbDatabaseName,
		MongoDbCollectionName: feat.MongoDbCollectionName,
	}

	var logInHandler *LogInHandler = &LogInHandler{
		JwtSecret:            feat.JwtSecret,
		GithubTokenValidator: githubTokenValidator,
		LogInDbService:       logInDbService,
	}

	return logInHandler
}
