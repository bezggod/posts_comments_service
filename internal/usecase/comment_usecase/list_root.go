package comment_usecase

import "posts_commets_service/internal/domain/models"

func (u *CommentUsecase) ListRoots(postID models.PostID, limit int, lastID *models.CommentID) ([]*models.Comment, *models.CommentID, error) {
	if limit < 0 {
		limit = 0
	}
	return u.comments.ListRoots(postID, limit, lastID)
}
