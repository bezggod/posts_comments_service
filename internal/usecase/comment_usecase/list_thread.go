package comment_usecase

import "posts_commets_service/internal/domain/models"

func (u *CommentUsecase) ListThread(postID models.PostID, firstCommentID models.CommentID, limit int, lastID *models.CommentID) ([]*models.Comment, *models.CommentID, error) {
	if limit < 0 {
		limit = 0
	}
	return u.comments.ListThreads(postID, firstCommentID, limit, lastID)
}
