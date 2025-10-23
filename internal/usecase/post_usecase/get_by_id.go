package post_usecase

import (
	"context"
	"posts_comments_service/internal/domain/models"
)

// GetByID возвращает пост по его id
func (u *PostUseCase) GetByID(ctx context.Context, id models.PostID) (*models.Post, error) {
	return u.posts.GetByID(ctx, id)
}
