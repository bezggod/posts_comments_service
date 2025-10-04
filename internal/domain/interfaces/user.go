package interfaces

import "posts_commets_service/internal/domain/models"

type UserRepo interface {
	Create(name string) (*models.User, error)
	GetByID(id models.UserID) (*models.User, error)
}
