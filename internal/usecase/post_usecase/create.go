package post_usecase

import (
	"context"
	"errors"
	"fmt"
	"posts_comments_service/internal/domain/models"
	"strings"
)

// Create создает новый пост
func (u *PostUseCase) Create(ctx context.Context, userID models.UserID, title, body string, commentBlock bool) (*models.Post, error) {
	if strings.TrimSpace(title) == "" {
		return nil, errors.New("title is empty")
	}
	if strings.TrimSpace(body) == "" {
		return nil, errors.New("body is empty")
	}
	post := &models.Post{
		UserID:       userID,
		Title:        title,
		Body:         body,
		CommentBlock: commentBlock,
	}

	createdPost, err := u.posts.Create(ctx, post)
	if err != nil {
		return nil, fmt.Errorf("PostUseCase.Create: %w", err)
	}
	return createdPost, nil
}
