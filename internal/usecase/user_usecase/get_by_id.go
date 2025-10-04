package user_usecase

import "posts_commets_service/internal/domain/models"

func (u *UserUseCase) GetByID(id models.UserID) (*models.User, error) {
	return u.users.GetByID(id)
}
