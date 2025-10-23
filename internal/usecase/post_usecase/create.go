package post_usecase

import (
	"context"
	"posts_comments_service/internal/domain/models"
)

// Create создает новый пост
func (u *PostUseCase) Create(ctx context.Context, userID models.UserID, title, body string, commentBlock bool) (*models.Post, error) {
	p, err := u.posts.Create(ctx, userID, title, body, commentBlock)
	if err != nil {
		return nil, err
	}
	return p, nil
}
