package user_storage

import (
	"context"
	"errors"
	"posts_commets_service/internal/domain/models"
)

func (r *UserRepo) GetByID(ctx context.Context, id models.UserID) (*models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, ok := r.byID[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}
