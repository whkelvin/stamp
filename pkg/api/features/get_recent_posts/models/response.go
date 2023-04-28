package models

import (
	"time"
)

type Response struct {
	Count    int    `json:"count"`
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	Posts    []Post `json:"posts"`
}

type Post struct {
	Id          string    `json:"id"`
	Link        string    `json:"link"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedDate time.Time `json:"createdDate"`
	RootDomain  string    `json:"rootDomain"`
}
