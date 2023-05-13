package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type createUserSchema struct {
	IsAdmin       bool           `bson:"isAdmin"`
	AuthProviders []authProvider `bson:"authProviders"`
	CreatedDate   string         `bson:"createdDate"`
}

type userSchema struct {
	Id            primitive.ObjectID `bson:"_id"`
	IsAdmin       bool               `bson:"isAdmin"`
	AuthProviders []authProvider     `bson:"authProviders"`
	CreatedDate   string             `bson:"createdDate"`
}

type authProvider struct {
	Name     string `bson:"name"`
	Username string `bson:"username"`
}
