package models

import "time"

type Response struct {
	PostId      string    `db:"post_id"`
	Link        string    `db:"link"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	CreatedDate time.Time `db:"created_date"`
	RootDomain  string    `db:"root_domain"`
}
