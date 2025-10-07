package post_usecase

import (
	"context"
	"posts_commets_service/internal/domain/models"
)

const (
	defaultLimit = 20
	maxLimit     = 100
)

func (u *PostUseCase) ListPosts(ctx context.Context, limit int, lastID *models.PostID) ([]*models.Post, *models.PostID, error) {
	if limit < 0 {
		limit = defaultLimit
	} else if limit > maxLimit {
		limit = maxLimit
	}
	return u.posts.ListPosts(ctx, limit, lastID)
}
