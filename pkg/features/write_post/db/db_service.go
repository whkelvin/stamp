package db

import (
	"context"
	"github.com/labstack/gommon/log"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/whkelvin/stamp/pkg/features/write_post/db/models"
)

type IWritePostDbService interface {
	CreatePost(ctx context.Context, newPost models.NewPost) (*models.PostDto, error)
}

type WritePostDbService struct {
	ConnPool *pgxpool.Pool
}

func (db *WritePostDbService) CreatePost(ctx context.Context, newPost models.NewPost) (*models.PostDto, error) {
	uuid := uuid.New()

	_, err := db.ConnPool.Exec(ctx, "insert into posts (post_id, link, title, description, created_date, root_domain) VALUES ($1, $2, $3, $4, now() at time zone('utc'), $5);", uuid, newPost.Link, newPost.Title, newPost.Description, newPost.RootDomain)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	var postDto models.PostDto
	rows, err := db.ConnPool.Query(ctx, "select post_id, link, title, description, created_date, root_domain from posts where post_id=$1", uuid)
	if err != nil {
		log.Error(err)
		return nil, err

	}

	err = pgxscan.ScanOne(&postDto, rows)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &postDto, nil
}
