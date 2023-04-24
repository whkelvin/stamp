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
	PostId      string
	Link        string
	Title       string
	Description string
	CreatedDate time.Time
	RootDomain  string
}
