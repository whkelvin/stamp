package get_recent_posts

import (
	"github.com/jackc/pgx/v5/pgxpool"
	. "github.com/whkelvin/stamp/pkg/features/get_recent_posts/db"
	. "github.com/whkelvin/stamp/pkg/features/get_recent_posts/handler"
)

type GetRecentPostsFeature struct {
	ConnPool *pgxpool.Pool
}

func (feat *GetRecentPostsFeature) Init() *GetRecentPostsHandler {
	var getRecentPostsDbService *GetRecentPostsDbService = &GetRecentPostsDbService{ConnPool: feat.ConnPool}
	var getRecentPostsHandler *GetRecentPostsHandler = &GetRecentPostsHandler{DbService: getRecentPostsDbService}
	return getRecentPostsHandler
}
