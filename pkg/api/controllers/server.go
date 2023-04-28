package controllers

import (
	get_recent_posts_handler "github.com/whkelvin/stamp/pkg/features/get_recent_posts/handler"
	write_post_handler "github.com/whkelvin/stamp/pkg/features/write_post/handler"
)

type ApiServer struct {
	GetRecentPostsHandler get_recent_posts_handler.IGetRecentPostsHandler
	WritePostHandler      write_post_handler.IWritePostHandler
}
