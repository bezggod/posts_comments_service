package user_usecase

import (
	"posts_commets_service/internal/usecase/user_usecase/mocks"
	"testing"
)

type mockService struct {
	userRepo *mocks.UserRepo
}

func makeService(t *testing.T) (*UserUseCase, mockService) {
	m := mockService{
		userRepo: mocks.NewUserRepo(t),
	}

	u := &UserUseCase{
		users: m.userRepo,
	}
	return u, m
}
