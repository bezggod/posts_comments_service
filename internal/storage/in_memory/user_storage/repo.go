package user_storage

import (
	"posts_commets_service/internal/domain/models"
	"sync"
	"sync/atomic"
)

type UserRepo struct {
	mu     sync.RWMutex
	nextID int64
	byID   map[models.UserID]*models.User
}

func NewUserRepo() *UserRepo {
	return &UserRepo{
		byID: make(map[models.UserID]*models.User),
	}
}
func (r *UserRepo) next() models.UserID {
	return models.UserID(atomic.AddInt64(&r.nextID, 1))
}
