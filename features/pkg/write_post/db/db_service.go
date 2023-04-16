package db

import (
	// "go_api_starter/database/pkg/db"
	// "go_api_starter/database/pkg/service"
	// "go_api_starter/features/pkg/get_user/db/models"
	//. "github.com/whkelvin/stamp/features/pkg/write_post/db"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/whkelvin/stamp/features/pkg/write_post/db/models"
)

type IWritePostDbService interface {
	CreatePost(models.NewPost) error
}

type WritePostDbService struct {
	Database *pgx.Conn
}

func (db *WritePostDbService) CreatePost(newPost models.NewPost) error {
	_, err := db.Database.Exec(context.Background(), "INSERT INTO post (link, title, description) VALUES ($1, $2, $3);", newPost.Link, newPost.Title, newPost.Description)
	if err != nil {
		return err
	}
	return nil
}
