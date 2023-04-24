package write_post

import (
	"github.com/jackc/pgx/v5/pgxpool"
	. "github.com/whkelvin/stamp/pkg/features/write_post/db"
	. "github.com/whkelvin/stamp/pkg/features/write_post/handler"
)

type WritePostFeature struct {
	ConnPool *pgxpool.Pool
}

func (feat *WritePostFeature) Init() *WritePostHandler {
	var writePostDbService *WritePostDbService = &WritePostDbService{ConnPool: feat.ConnPool}
	var writePostHandler *WritePostHandler = &WritePostHandler{DbService: writePostDbService}
	return writePostHandler
}
