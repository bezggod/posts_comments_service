package interfaces

import (
	"context"
	"posts_commets_service/internal/domain/models"
)

type CommentRepo interface {
	CreateReplyComment(ctx context.Context, postID models.PostID, userID models.UserID, parentCommentID models.CommentID, text string) (*models.Comment, error)
	CreateRoot(ctx context.Context, postID models.PostID, userID models.UserID, text string) (*models.Comment, error)
	GetByID(ctx context.Context, id models.CommentID) (*models.Comment, error)
	ListRoots(ctx context.Context, postID models.PostID, limit int, lastID *models.CommentID) ([]*models.Comment, *models.CommentID, error)
	ListThreads(ctx context.Context, postID models.PostID, firstCommentID models.CommentID, limit int, lastID *models.CommentID) ([]*models.Comment, *models.CommentID, error)
}
