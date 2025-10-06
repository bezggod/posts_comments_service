package comment_storage

import (
	"context"
	"fmt"
	"posts_commets_service/internal/domain/models"
	"time"
)

func (r *CommentRepo) CreateRoot(ctx context.Context, postID models.PostID, userID models.UserID, text string) (*models.Comment, error) {
	if len(text) > maxLength {
		return nil, fmt.Errorf("text very long")
	}

	id := r.next()
	c := &models.Comment{
		ID:              id,
		PostID:          postID,
		UserID:          userID,
		Text:            text,
		ParentCommentID: nil,
		FirstCommentID:  id,
		CreatedAt:       time.Now(),
	}
	r.mu.Lock()
	r.byID[id] = c
	r.mu.Unlock()
	return c, nil

}
