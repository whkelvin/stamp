package auth

import (
	"github.com/golang-jwt/jwt/v4"
)

type JwtClaims struct {
	Username         string `json:"username"`
	AuthProviderName string `json:"authProviderName"`
	IsAdmin          bool   `json:"isAdmin"`
	jwt.RegisteredClaims
}
