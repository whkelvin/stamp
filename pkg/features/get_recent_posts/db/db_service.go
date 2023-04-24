package db

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/gommon/log"
	. "github.com/whkelvin/stamp/pkg/features/get_recent_posts/db/models"
)

type IGetRecentPostsDbService interface {
	GetRecentPosts(ctx context.Context, req Request) (*Response, error)
}

type GetRecentPostsDbService struct {
	ConnPool *pgxpool.Pool
}

func (db *GetRecentPostsDbService) GetRecentPosts(ctx context.Context, req Request) (*Response, error) {

	tx, err := db.ConnPool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel:       pgx.ReadCommitted,
		AccessMode:     pgx.ReadOnly,
		DeferrableMode: pgx.Deferrable,
	})
	if err != nil {
		return nil, err
	}

	defer tx.Rollback(ctx)

	var posts []Post
	rows, err := db.ConnPool.Query(ctx, "select post_id, link, title, description, created_date, root_domain from posts order by created_date desc limit $1 offset $2", req.Take, req.Skip)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	err = pgxscan.ScanAll(&posts, rows)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	var count int
	row := db.ConnPool.QueryRow(ctx, "SELECT COUNT(*) FROM posts")
	err = row.Scan(&count)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if err = tx.Commit(ctx); err != nil {
		panic(err)
	}

	res := &Response{
		Posts:      posts,
		Count:      len(posts),
		TotalCount: count,
	}

	return res, nil
}
