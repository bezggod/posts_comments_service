package service_provider

import (
	"posts_commets_service/internal/config"
	"posts_commets_service/internal/controller/http_controller/comment"
	"posts_commets_service/internal/controller/http_controller/post"
	"posts_commets_service/internal/controller/http_controller/user"
	"posts_commets_service/internal/domain/interfaces"
	"posts_commets_service/internal/usecase/comment_usecase"
	"posts_commets_service/internal/usecase/post_usecase"
	"posts_commets_service/internal/usecase/user_usecase"
)

type ServiceProvider struct {
	userRepo    interfaces.UserRepo
	postRepo    interfaces.PostRepo
	commentRepo interfaces.CommentRepo

	userUC    *user_usecase.UserUseCase
	postUC    *post_usecase.PostUseCase
	commentUC *comment_usecase.CommentUseCase

	userController    *user.Controller
	postController    *post.Controller
	commentController *comment.Controller

	dbCluster *config.Cluster
}

func NewServiceProvider(storage string) *ServiceProvider {
	return &ServiceProvider{}
}
