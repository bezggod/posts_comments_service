package service_provider

import (
	"posts_commets_service/internal/controller/http_controller/comment"
	"posts_commets_service/internal/controller/http_controller/post"
	"posts_commets_service/internal/controller/http_controller/user"
)

func (s *ServiceProvider) GetUserController() *user.Controller {
	if s.userController == nil {
		s.userController = user.NewUserController(s.GetUserUseCase())
	}
	return s.userController
}

func (s *ServiceProvider) GetPostController() *post.Controller {
	if s.postController == nil {
		s.postController = post.NewPostController(s.GetPostUseCase())
	}
	return s.postController
}

func (s *ServiceProvider) GetCommentController() *comment.Controller {
	if s.commentController == nil {
		s.commentController = comment.NewCommentController(s.GetCommentUseCase())
	}
	return s.commentController
}
