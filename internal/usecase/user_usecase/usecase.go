package user_usecase

import (
	"posts_comments_service/internal/domain/interfaces"
)

type UserUseCase struct {
	users interfaces.UserRepo
}

type PostUseCase struct {
	posts interfaces.PostRepo
}

func NewUserUseCase(users interfaces.UserRepo) *UserUseCase {
	return &UserUseCase{users: users}
}

func NewPostUseCase(posts interfaces.PostRepo) *PostUseCase {
	return &PostUseCase{posts: posts}
}
