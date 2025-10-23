package comment_usecase

import (
	"github.com/stretchr/testify/assert"
	"posts_comments_service/internal/domain/interfaces/mocks"
	"testing"
)

type mockService struct {
	postRepo    *mocks.MockPostRepo
	commentRepo *mocks.MockCommentRepo
}

func TestCommentUsecase(t *testing.T) {
	u := NewCommentUseCase(nil, nil)
	assert.NotNil(t, u)
}

func makeService(t *testing.T) (*CommentUseCase, mockService) {
	m := mockService{
		postRepo:    mocks.NewMockPostRepo(t),
		commentRepo: mocks.NewMockCommentRepo(t),
	}

	u := &CommentUseCase{
		comments: m.commentRepo,
		posts:    m.postRepo,
	}
	return u, m
}
