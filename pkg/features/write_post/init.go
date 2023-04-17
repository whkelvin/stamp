package write_post

import (
	"github.com/jackc/pgx/v5"
	. "github.com/whkelvin/stamp/pkg/features/write_post/db"
	. "github.com/whkelvin/stamp/pkg/features/write_post/handler"
)

type WritePostFeature struct {
	Database *pgx.Conn
}

func (feat *WritePostFeature) Init() *WritePostHandler {
	var writePostDbService *WritePostDbService = &WritePostDbService{Database: feat.Database}
	var writePostHandler *WritePostHandler = &WritePostHandler{DbService: writePostDbService}
	return writePostHandler
}
