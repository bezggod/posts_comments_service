package post_storage

import (
	"fmt"
	"posts_commets_service/internal/domain/models"
)

func (r *PostRepo) GetById(id models.PostID) (*models.Post, error) {
	r.mu.RLock()
	p := r.byID[id]
	r.mu.RUnlock()
	if p == nil {
		return nil, fmt.Errorf("post not found")
	}
	return p, nil
}
