package interfaces

import (
	"context"
	"posts_comments_service/internal/domain/models"
)

type UserRepo interface {
	Create(ctx context.Context, name string) (*models.User, error)
	GetByID(ctx context.Context, id models.UserID) (*models.User, error)
}
