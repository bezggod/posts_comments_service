package comment_storage

import (
	"context"
	"fmt"
	"posts_commets_service/internal/domain/models"
	"time"
)

func (r *CommentRepo) CreateReplyComment(ctx context.Context, postID models.PostID, userID models.UserID, parentCommentID models.CommentID, text string) (*models.Comment, error) {
	if len(text) > maxLength {
		return nil, fmt.Errorf("comment text is very long")
	}

	r.mu.RLock()
	parent := r.byID[parentCommentID]
	r.mu.RUnlock()
	if parent == nil {
		return nil, fmt.Errorf("parent comment not found")
	}
	if parent.PostID != postID {
		return nil, fmt.Errorf("parent comment does not belong to post")
	}

	id := r.next()
	c := &models.Comment{
		ID:              id,
		PostID:          postID,
		UserID:          userID,
		Text:            text,
		ParentCommentID: &parentCommentID,
		FirstCommentID:  parent.FirstCommentID,
		CreatedAt:       time.Now(),
	}
	r.mu.Lock()
	r.byID[id] = c
	r.mu.Unlock()
	return c, nil
}
