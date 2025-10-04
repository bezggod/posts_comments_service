package comment_storage

import (
	"posts_commets_service/internal/domain/models"
	"sync"
	"sync/atomic"
)

const maxLength = 2000

type CommentRepo struct {
	mu     sync.RWMutex
	nextID int64
	byID   map[models.CommentID]*models.Comment
}

func NewCommentRepo() *CommentRepo {
	return &CommentRepo{
		byID: make(map[models.CommentID]*models.Comment),
	}
}

func (r *CommentRepo) next() models.CommentID {
	return models.CommentID(atomic.AddInt64(&r.nextID, 1))
}
