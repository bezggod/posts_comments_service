package post

import "posts_comments_service/internal/usecase/post_usecase"

type Controller struct {
	uc *post_usecase.PostUseCase
}

func NewPostController(uc *post_usecase.PostUseCase) *Controller {
	return &Controller{
		uc: uc,
	}
}
