package interfaces

import "posts_commets_service/internal/domain/models"

type PostRepo interface {
	Create(userID models.UserID, title, body string, commentBlock bool) (*models.Post, error)
	GetById(id models.PostID) (*models.Post, error)
	ListPosts(limit int, lastID *models.PostID) ([]*models.Post, *models.PostID, error)
	CommentBlock(id models.PostID, off bool) error
}
