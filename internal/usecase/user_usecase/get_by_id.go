package user_usecase

import (
	"context"
	"posts_comments_service/internal/domain/models"
)

// GetByID возвращает пользователя по id
func (u *UserUseCase) GetByID(ctx context.Context, id models.UserID) (*models.User, error) {
	return u.users.GetByID(ctx, id)
}
