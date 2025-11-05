package user_usecase

import (
	"context"
	"fmt"
	"posts_comments_service/internal/domain/models"
)

type UseCase struct {
	Name string
}

// Create создает нового пользователя
func (u *UserUseCase) Create(ctx context.Context, name string) (*models.User, error) {
	user := &models.User{
		Name: name,
	}
	createdUser, err := u.users.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("UserUseCase.Create: %w", err)
	}
	return createdUser, nil
}
