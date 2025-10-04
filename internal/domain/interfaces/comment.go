package interfaces

import "posts_commets_service/internal/domain/models"

type CommentRepo interface {
	CreateReplyComment(postID models.PostID, userID models.UserID, parentCommentID models.CommentID, text string) (*models.Comment, error)
	CreateRoot(postID models.PostID, userID models.UserID, text string) (*models.Comment, error)
	GetByID(id models.CommentID) (*models.Comment, error)
	ListRoots(postID models.PostID, limit int, lastID *models.CommentID) ([]*models.Comment, *models.CommentID, error)
	ListThreads(postID models.PostID, firstCommentID models.CommentID, limit int, lastID *models.CommentID) ([]*models.Comment, *models.CommentID, error)
}
