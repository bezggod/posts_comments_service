package post_storage

import (
	"fmt"
	"posts_commets_service/internal/domain/models"
	"time"
)

func (r *PostRepo) Create(userID models.UserID, title, body string, commentBlock bool) (*models.Post, error) {
	if title == "" {
		return nil, fmt.Errorf("title is empty")
	}
	id := r.next()
	p := &models.Post{
		ID:           id,
		UserID:       userID,
		Title:        title,
		Body:         body,
		CommentBlock: commentBlock,
		CreatedAt:    time.Now(),
	}

	r.mu.Lock()
	r.byID[id] = p
	r.mu.Unlock()
	return p, nil
}
