package post_usecase

import (
	"context"
	"posts_commets_service/internal/domain/models"
)

func (u *PostUseCase) GetByID(ctx context.Context, id models.PostID) (*models.Post, error) {
	return u.posts.GetByID(ctx, id)
}
