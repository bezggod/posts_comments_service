package comment_usecase

import "posts_commets_service/internal/domain/interfaces"

type CommentUsecase struct {
	posts    interfaces.PostRepo
	comments interfaces.CommentRepo
}

func NewCommentUseCase(posts interfaces.PostRepo, comments interfaces.CommentRepo) *CommentUsecase {
	return &CommentUsecase{
		posts:    posts,
		comments: comments,
	}
}
