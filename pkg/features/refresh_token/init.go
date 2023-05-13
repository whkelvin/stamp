package refresh_token

import (
	. "github.com/whkelvin/stamp/pkg/features/refresh_token/handler"
)

type RefreshTokenFeature struct {
	JwtSecret string
}

func (feat *RefreshTokenFeature) Init() *RefreshTokenHandler {
	var refreshTokenHandler *RefreshTokenHandler = &RefreshTokenHandler{
		JwtSecret: feat.JwtSecret,
	}

	return refreshTokenHandler
}
