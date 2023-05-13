package handler

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	. "github.com/whkelvin/stamp/pkg/auth"
	. "github.com/whkelvin/stamp/pkg/features/refresh_token/handler"
	handlerModels "github.com/whkelvin/stamp/pkg/features/refresh_token/handler/models"
	"testing"
	"time"
)

func Test_Handler_Should_Return_Token_When_Input_Is_Valid(t *testing.T) {

	var handler IRefreshTokenHandler = &RefreshTokenHandler{JwtSecret: "secret"}

	claims := &JwtClaims{
		Username:         "test",
		IsAdmin:          true,
		AuthProviderName: "github",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign token with secret
	signedToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		assert.Fail(t, "Cannot sign token")
	}

	res, err := handler.RefreshToken(context.Background(), handlerModels.Request{Jwt: signedToken})

	if err != nil {
		assert.Fail(t, "Refresh Token Failed")
	}

	if res.Jwt == "" {
		assert.Fail(t, "Returned token is empty")
	}
}
