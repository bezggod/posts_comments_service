package comment_usecase

import (
	"context"
	"posts_comments_service/internal/domain/models"
)

// ListRoots возвращает список корневых комментариев для поста
func (u *CommentUseCase) ListRoots(ctx context.Context, postID models.PostID, limit int, lastID *models.CommentID) ([]*models.Comment, *models.CommentID, error) {

	if limit < 0 {
		limit = 0
	}
	return u.comments.ListRoots(ctx, postID, limit, lastID)
}
