package post_usecase

import (
	"context"
	"errors"
	"posts_commets_service/internal/domain/models"
)

func (u *PostUseCase) GetByID(ctx context.Context, id models.PostID) (*models.Post, error) {
	if id == 0 {
		return nil, errors.New("id empty")
	}
	return u.posts.GetByID(ctx, id)
}
