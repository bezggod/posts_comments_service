package user_usecase

import (
	"context"
	"posts_comments_service/internal/domain/models"
)

type CreateUserReq struct {
	Name string
}

// Create создает нового пользователя
func (u *UserUseCase) Create(ctx context.Context, name string) (*models.User, error) {
	return u.users.Create(ctx, name)
}
