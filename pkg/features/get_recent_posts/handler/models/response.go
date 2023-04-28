package models

import (
	"time"
)

type Response struct {
	Count int
	Posts []Post
}

type Post struct {
	Id          string
	Link        string
	Title       string
	Description string
	CreatedDate time.Time
	RootDomain  string
}
