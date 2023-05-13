package handler

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	. "github.com/whkelvin/stamp/pkg/auth"
	handlerError "github.com/whkelvin/stamp/pkg/features/errors/handler"
	"github.com/whkelvin/stamp/pkg/features/log_in/db"
	dbModels "github.com/whkelvin/stamp/pkg/features/log_in/db/models"
	handlerModels "github.com/whkelvin/stamp/pkg/features/log_in/handler/models"
	"github.com/whkelvin/stamp/pkg/features/log_in/helpers"
	"time"
)

type ILogInHandler interface {
	LogIn(ctx context.Context, req handlerModels.Request) (*handlerModels.Response, error)
}

type LogInHandler struct {
	JwtSecret            string
	GithubTokenValidator helpers.IGithubTokenValidator
	LogInDbService       db.ILogInDbService
}

func (handler *LogInHandler) LogIn(ctx context.Context, req handlerModels.Request) (*handlerModels.Response, error) {

	if req.AuthProvider == "github" {
		ghUser, err := handler.GithubTokenValidator.ValidateGithubToken(req.AccessToken)
		if err != nil {
			return nil, handlerError.New(err.Error(), true)
		}

		var dbReq dbModels.Request = dbModels.Request{
			Username:         ghUser.Username,
			AuthProviderName: "github",
		}
		user, err := handler.LogInDbService.CreateOrGetUser(ctx, dbReq)
		if err != nil {
			return nil, handlerError.New(err.Error(), false)
		}
		if user == nil {
			return nil, handlerError.New("Unable to log in or sign up", false)
		}

		claims := &JwtClaims{
			Username:         user.Username,
			IsAdmin:          user.IsAdmin,
			AuthProviderName: user.AuthProviderName,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
			},
		}

		// Create token with claims
		t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		// Sign token with secret
		token, err := t.SignedString([]byte(handler.JwtSecret))
		if err != nil {
			return nil, handlerError.New(err.Error(), false)
		}
		return &handlerModels.Response{
			JwtToken: token,
		}, nil
	}

	return nil, handlerError.New("token not supported", true)

}
