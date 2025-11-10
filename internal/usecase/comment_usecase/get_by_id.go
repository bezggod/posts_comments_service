package comment_usecase

import (
	"context"
	"posts_comments_service/internal/domain/models"
)

// GetByID возвращает комментарий по его ID
func (u *CommentUseCase) GetByID(ctx context.Context, id models.CommentID) (*models.Comment, error) {
	return u.comments.GetByID(ctx, id)
}
