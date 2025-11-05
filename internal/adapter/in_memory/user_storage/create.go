package user_storage

import (
	"context"
	"fmt"
	"posts_comments_service/internal/domain/models"
)

func (r *UserRepo) Create(ctx context.Context, user *models.User) (*models.User, error) {
	if user == nil {
		return nil, fmt.Errorf("user is nil")
	}

	if user.Name == "" {
		return nil, fmt.Errorf("name is empty")
	}
	id := r.next()
	n := &models.User{
		ID:   id,
		Name: user.Name,
	}

	r.mu.Lock()
	r.byID[id] = n
	r.mu.Unlock()
	return n, nil
}
