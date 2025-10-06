package user_storage

import (
	"context"
	"fmt"
	"posts_commets_service/internal/domain/models"
)

func (r *UserRepo) Create(ctx context.Context, name string) (*models.User, error) {
	if name == "" {
		return nil, fmt.Errorf("name is empty")
	}
	id := r.next()
	n := &models.User{
		ID:   id,
		Name: name,
	}

	r.mu.Lock()
	r.byID[id] = n
	r.mu.Unlock()
	return n, nil
}
