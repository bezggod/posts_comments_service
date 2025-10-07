package graphql

import (
	"posts_commets_service/internal/usecase/comment_usecase"
	"posts_commets_service/internal/usecase/post_usecase"
	"posts_commets_service/internal/usecase/user_usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CommentUC *comment_usecase.CommentUseCase
	PostUC    *post_usecase.PostUseCase
	UserUC    *user_usecase.UserUseCase
}
