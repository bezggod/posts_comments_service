package comment_usecase

import (
	"context"
	"posts_commets_service/internal/domain/models"
)

func (u *CommentUseCase) GetByID(ctx context.Context, id models.CommentID) (*models.Comment, error) {
	return u.comments.GetByID(ctx, id)
}
