package comment_usecase

import "posts_commets_service/internal/domain/interfaces"

type CommentUseCase struct {
	posts    interfaces.PostRepo
	comments interfaces.CommentRepo
}

func NewCommentUseCase(posts interfaces.PostRepo, comments interfaces.CommentRepo) *CommentUseCase {
	return &CommentUseCase{
		posts:    posts,
		comments: comments,
	}
}
