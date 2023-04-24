package models

import (
	"time"
)

type Response struct {
	Count      int
	Posts      []Post
	TotalCount int
}

type Post struct {
	PostId      string    `db:"post_id"`
	Link        string    `db:"link"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	CreatedDate time.Time `db:"created_date"`
	RootDomain  string    `db:"root_domain"`
}
