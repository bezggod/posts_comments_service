package user

import "posts_commets_service/internal/usecase/user_usecase"

type Controller struct {
	uc *user_usecase.UserUseCase
}

func NewUserController(uc *user_usecase.UserUseCase) *Controller {
	return &Controller{
		uc: uc,
	}
}
