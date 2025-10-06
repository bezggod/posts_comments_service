package comment_usecase

import (
	"context"
	"errors"
	"posts_commets_service/internal/domain/models"
)

func (u *CommentUseCase) ListRoots(ctx context.Context, postID models.PostID, limit int, lastID *models.CommentID) ([]*models.Comment, *models.CommentID, error) {
	if postID == 0 {
		return nil, nil, errors.New("postID emprty")
	}

	if limit < 0 {
		limit = 0
	}
	return u.comments.ListRoots(ctx, postID, limit, lastID)
}
