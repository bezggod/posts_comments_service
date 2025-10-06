package comment_usecase

import (
	"context"
	"errors"
	"posts_commets_service/internal/domain/models"
)

func (u *CommentUseCase) ListThreads(ctx context.Context, postID models.PostID, firstCommentID models.CommentID, limit int, lastID *models.CommentID) ([]*models.Comment, *models.CommentID, error) {
	if postID == 0 || firstCommentID == 0 {
		return nil, nil, errors.New("invalid postID and firstCommentID")
	}
	if limit < 0 {
		limit = 0
	}
	return u.comments.ListThreads(ctx, postID, firstCommentID, limit, lastID)
}
