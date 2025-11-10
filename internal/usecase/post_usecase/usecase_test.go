package post_usecase

import (
	"github.com/stretchr/testify/assert"
	"posts_comments_service/internal/domain/interfaces/mocks"
	"testing"
)

type mockService struct {
	postRepo *mocks.MockPostRepo
}

func TestPostUseCase(t *testing.T) {
	u := NewPostUseCase(nil)
	assert.NotNil(t, u)
}

func makeService(t *testing.T) (*PostUseCase, mockService) {
	m := mockService{
		postRepo: mocks.NewMockPostRepo(t),
	}

	u := &PostUseCase{
		posts: m.postRepo,
	}
	return u, m
}
