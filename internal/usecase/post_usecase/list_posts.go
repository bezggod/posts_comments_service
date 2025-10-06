package post_usecase

import (
	"context"
	"posts_commets_service/internal/domain/models"
)

func (u *PostUseCase) ListPosts(ctx context.Context, limit int, lastID *models.PostID) ([]*models.Post, *models.PostID, error) {
	if limit < 0 {
		limit = 0
	}
	return u.posts.ListPosts(ctx, limit, lastID)
}
