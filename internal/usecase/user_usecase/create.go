package user_usecase

import (
	"errors"
	"posts_commets_service/internal/domain/models"
)

func (u *UserUseCase) Create(name string) (*models.User, error) {
	if name == "" {
		return nil, errors.New("name is empty")
	}
	return u.users.Create(name)
}
