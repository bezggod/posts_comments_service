package service_provider

import (
	"posts_commets_service/internal/usecase/comment_usecase"
	"posts_commets_service/internal/usecase/post_usecase"
	"posts_commets_service/internal/usecase/user_usecase"
)

func (s *ServiceProvider) GetUserUseCase() *user_usecase.UserUseCase {
	if s.userUC == nil {
		s.userUC = user_usecase.NewUserUseCase(s.getUserRepo())
	}
	return s.userUC
}

func (s *ServiceProvider) GetPostUseCase() *post_usecase.PostUseCase {
	if s.postUC == nil {
		s.postUC = post_usecase.NewPostUseCase(s.getPostRepo())
	}
	return s.postUC
}

func (s *ServiceProvider) GetCommentUseCase() *comment_usecase.CommentUseCase {
	if s.commentUC == nil {
		s.commentUC = comment_usecase.NewCommentUseCase(s.getPostRepo(), s.getCommentRepo())
	}
	return s.commentUC

}
