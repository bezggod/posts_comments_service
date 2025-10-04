package post_storage

import (
	"posts_commets_service/internal/domain/models"
	"sync"
	"sync/atomic"
)

type PostRepo struct {
	mu     sync.RWMutex
	nextID int64
	byID   map[models.PostID]*models.Post
}

func NewPostRepo() *PostRepo {
	return &PostRepo{
		byID: make(map[models.PostID]*models.Post),
	}
}
func (r *PostRepo) next() models.PostID {
	return models.PostID(atomic.AddInt64(&r.nextID, 1))
}
