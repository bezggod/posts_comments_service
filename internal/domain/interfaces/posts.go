package interfaces

import (
	"context"
	"posts_commets_service/internal/domain/models"
)

type PostRepo interface {
	Create(ctx context.Context, userID models.UserID, title, body string, commentBlock bool) (*models.Post, error)
	GetByID(ctx context.Context, id models.PostID) (*models.Post, error)
	ListPosts(ctx context.Context, limit int, lastID *models.PostID) ([]*models.Post, *models.PostID, error)
	CommentBlock(ctx context.Context, id models.PostID, off bool) error
}
