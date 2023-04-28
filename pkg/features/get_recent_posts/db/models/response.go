package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Response struct {
	Count int
	Posts []Post
}

type Post struct {
	Id          primitive.ObjectID `bson:"_id"`
	Link        string             `bson:"link"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	CreatedDate time.Time          `bson:"createdDate"`
	RootDomain  string             `bson:"rootDomain"`
}
