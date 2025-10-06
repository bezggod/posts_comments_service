package user_usecase

import (
	"context"
	"posts_commets_service/internal/domain/models"
)

func (u *UserUseCase) GetByID(ctx context.Context, id models.UserID) (*models.User, error) {
	return u.users.GetByID(ctx, id)
}
