package user_usecase

import (
	"context"
	"errors"
	"posts_commets_service/internal/domain/models"
)

type CreateUserReq struct {
	Name string
}

func (u *UserUseCase) Create(ctx context.Context, name string) (*models.User, error) {
	if name == "" {
		return nil, errors.New("name is empty")
	}
	return u.users.Create(ctx, name)
}
