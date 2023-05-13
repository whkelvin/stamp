package controllers

import (
	get_recent_posts_handler "github.com/whkelvin/stamp/pkg/features/get_recent_posts/handler"
	log_in_handler "github.com/whkelvin/stamp/pkg/features/log_in/handler"
	refresh_token_handler "github.com/whkelvin/stamp/pkg/features/refresh_token/handler"
	write_post_handler "github.com/whkelvin/stamp/pkg/features/write_post/handler"
)

type ApiServer struct {
	GetRecentPostsHandler get_recent_posts_handler.IGetRecentPostsHandler
	WritePostHandler      write_post_handler.IWritePostHandler
	LogInHandler          log_in_handler.ILogInHandler
	RefreshTokenHandler   refresh_token_handler.IRefreshTokenHandler
}
