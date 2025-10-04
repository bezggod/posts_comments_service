package comment_storage

import (
	"errors"
	"posts_commets_service/internal/domain/models"
)

func (r *CommentRepo) GetByID(id models.CommentID) (*models.Comment, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	comment, ok := r.byID[id]
	if !ok {
		return nil, errors.New("comment not found")
	}
	return comment, nil
}
