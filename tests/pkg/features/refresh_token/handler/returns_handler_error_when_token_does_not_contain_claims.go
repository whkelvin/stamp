package handler

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	handlerError "github.com/whkelvin/stamp/pkg/features/errors/handler"
	. "github.com/whkelvin/stamp/pkg/features/refresh_token/handler"
	handlerModels "github.com/whkelvin/stamp/pkg/features/refresh_token/handler/models"
	"testing"
)

func Test_Handler_Should_Return_Error_When_Token_Does_Not_Contain_Claims(t *testing.T) {
	var handler IRefreshTokenHandler = &RefreshTokenHandler{JwtSecret: "secret"}

	token := jwt.New(jwt.SigningMethodES256)

	signedToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		assert.Fail(t, "Cannot sign token")
	}

	_, err = handler.RefreshToken(context.Background(), handlerModels.Request{Jwt: signedToken})

	handlerErr, ok := err.(handlerError.HandlerError)
	if !ok || !handlerErr.IsBadInput() {
		assert.Fail(t, "handler did not return Bad Input Error")
	}
}
