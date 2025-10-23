package user_usecase

import (
	"github.com/stretchr/testify/assert"
	"posts_comments_service/internal/domain/interfaces/mocks"
	"testing"
)

type mockService struct {
	userRepo *mocks.MockUserRepo
}

func TestUserUsecase(t *testing.T) {
	u := NewUserUseCase(nil)
	assert.NotNil(t, u)
}

func TestPostUsecase(t *testing.T) {
	u := NewPostUseCase(nil)
	assert.NotNil(t, u)
}

func makeService(t *testing.T) (*UserUseCase, mockService) {
	m := mockService{
		userRepo: mocks.NewMockUserRepo(t),
	}

	u := &UserUseCase{
		users: m.userRepo,
	}
	return u, m
}
