package post_usecase

import (
	"posts_commets_service/internal/usecase/post_usecase/mocks"
	"testing"
)

type mockService struct {
	postRepo *mocks.PostRepo
}

func makeService(t *testing.T) (*PostUseCase, mockService) {
	m := mockService{
		postRepo: mocks.NewPostRepo(t),
	}

	u := &PostUseCase{
		posts: m.postRepo,
	}
	return u, m
}
