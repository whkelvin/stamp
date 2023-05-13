package handler

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	. "github.com/whkelvin/stamp/pkg/auth"
	handlerError "github.com/whkelvin/stamp/pkg/features/errors/handler"
	handlerModels "github.com/whkelvin/stamp/pkg/features/refresh_token/handler/models"
	"time"
)

type IRefreshTokenHandler interface {
	RefreshToken(ctx context.Context, req handlerModels.Request) (*handlerModels.Response, error)
}

type RefreshTokenHandler struct {
	JwtSecret string
}

func (handler *RefreshTokenHandler) RefreshToken(ctx context.Context, req handlerModels.Request) (*handlerModels.Response, error) {
	claims := JwtClaims{}
	_, err := jwt.ParseWithClaims(req.Jwt, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(handler.JwtSecret), nil
	})

	if err != nil {
		return nil, handlerError.New(err.Error(), true)
	}

	newClaims := &JwtClaims{
		Username:         claims.Username,
		IsAdmin:          claims.IsAdmin,
		AuthProviderName: claims.AuthProviderName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	token, err := t.SignedString([]byte(handler.JwtSecret))
	if err != nil {
		return nil, handlerError.New(err.Error(), false)
	}
	return &handlerModels.Response{
		Jwt: token,
	}, nil
}
