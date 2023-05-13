package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	. "github.com/whkelvin/stamp/pkg/api/generated"

	//handlerError "github.com/whkelvin/stamp/pkg/features/errors/handler"
	"strings"

	handlerModels "github.com/whkelvin/stamp/pkg/features/refresh_token/handler/models"
)

func (server *ApiServer) RefreshToken(ctx echo.Context) error {
	rawToken := ctx.Request().Header.Get("Authorization")
	token := rawToken

	if strings.Contains(rawToken, "Bearer") {
		token = strings.TrimPrefix(rawToken, "Bearer ")
	} else if strings.Contains(rawToken, "bearer") {
		token = strings.TrimPrefix(rawToken, "bearer ")
	} else {
		ctx.String(http.StatusBadRequest, "Refresh Failed")
	}

	handlerRes, err := server.RefreshTokenHandler.RefreshToken(ctx.Request().Context(), handlerModels.Request{
		Jwt: token,
	})

	if err != nil || handlerRes == nil {
		ctx.String(http.StatusBadRequest, "Refresh Failed")
	}

	res := RefreshTokenResponse{
		Jwt: handlerRes.Jwt,
	}

	return ctx.JSON(http.StatusOK, res)
}
