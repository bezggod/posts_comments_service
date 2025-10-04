package user_usecase

import "posts_commets_service/internal/domain/interfaces"

type UserUseCase struct {
	users interfaces.UserRepo
}

func NewUserUseCase(users interfaces.UserRepo) *UserUseCase {
	return &UserUseCase{users: users}
}
