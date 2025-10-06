package comment_usecase

import (
	ctmocks "posts_commets_service/internal/usecase/comment_usecase/mocks"
	ptmocks "posts_commets_service/internal/usecase/post_usecase/mocks"
	"testing"
)

type mockService struct {
	postRepo    *ptmocks.PostRepo
	commentRepo *ctmocks.CommentRepo
}

func makeService(t *testing.T) (*CommentUseCase, mockService) {
	m := mockService{
		postRepo:    ptmocks.NewPostRepo(t),
		commentRepo: ctmocks.NewCommentRepo(t),
	}

	u := &CommentUseCase{
		comments: m.commentRepo,
		posts:    m.postRepo,
	}
	return u, m
}
