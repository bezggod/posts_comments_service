package post_storage

import (
	"context"
	"fmt"
	"posts_commets_service/internal/domain/models"
)

func (r *PostRepo) GetByID(ctx context.Context, id models.PostID) (*models.Post, error) {
	r.mu.RLock()
	p := r.byID[id]
	r.mu.RUnlock()
	if p == nil {
		return nil, fmt.Errorf("post not found")
	}
	return p, nil
}
