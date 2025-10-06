package post_usecase

import (
	"context"
	"fmt"
	"posts_commets_service/internal/domain/models"
)

func (u *PostUseCase) Create(ctx context.Context, userID models.UserID, title, body string, commentBlock bool) (*models.Post, error) {
	if title == "" {
		return nil, fmt.Errorf("title is empty")
	}
	if body == "" {
		return nil, fmt.Errorf("body is empty")
	}
	p, err := u.posts.Create(ctx, userID, title, body, commentBlock)
	if err != nil {
		return nil, err
	}
	return p, nil
}
