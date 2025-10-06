package post_storage

import (
	"context"
	"errors"
	"posts_commets_service/internal/domain/models"
)

func (r *PostRepo) CommentBlock(ctx context.Context, id models.PostID, off bool) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	p := r.byID[id]
	if p == nil {
		return errors.New("post not found")
	}
	p.CommentBlock = off
	return nil
}
