package controllers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	. "github.com/whkelvin/stamp/pkg/api/generated"
	handlerError "github.com/whkelvin/stamp/pkg/features/errors/handler"
	handlerModels "github.com/whkelvin/stamp/pkg/features/log_in/handler/models"
)

func parseLogInRequest(c echo.Context) (*LogInRequest, error) {
	var req LogInRequest
	err := c.Bind(&req)
	if err != nil {
		return nil, err
	}

	if req.AccessToken == "" {
		return nil, errors.New("field 'accessToken' is required.")
	}

	if req.AuthProvider == "" {
		return nil, errors.New("field 'authProvider' is required.")
	}

	if req.AuthProvider != "github" {
		return nil, errors.New(req.AuthProvider + " is currently not supported. Try one of the following: [github]")
	}

	return &req, nil
}

func (server *ApiServer) LogIn(ctx echo.Context) error {
	req, err := parseLogInRequest(ctx)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	handlerRes, err := server.LogInHandler.LogIn(ctx.Request().Context(), handlerModels.Request{
		AuthProvider: req.AuthProvider,
		AccessToken:  req.AccessToken,
	})
	if err != nil {
		handlerErr, ok := err.(handlerError.HandlerError)
		if ok && handlerErr.IsBadInput() {
			return ctx.String(http.StatusBadRequest, err.Error())
		}
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	res := LogInResponse{
		Jwt: handlerRes.JwtToken,
	}

	return ctx.JSON(http.StatusOK, res)
}
