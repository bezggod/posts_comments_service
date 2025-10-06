package comment

import "posts_commets_service/internal/usecase/comment_usecase"

type Controller struct {
	uc *comment_usecase.CommentUseCase
}

func NewCommentController(uc *comment_usecase.CommentUseCase) *Controller {
	return &Controller{
		uc: uc,
	}
}
