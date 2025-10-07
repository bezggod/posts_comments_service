package post_storage

import (
	"context"
	"posts_commets_service/internal/domain/models"
	"sort"
)

func (r *PostRepo) ListPosts(ctx context.Context, limit int, lastID *models.PostID) ([]*models.Post, *models.PostID, error) {
	r.mu.RLock()
	var posts []*models.Post
	for _, p := range r.byID {
		if lastID == nil || p.ID < *lastID {
			posts = append(posts, p)
		}
	}
	r.mu.RUnlock()

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].ID > posts[j].ID
	})

	if limit < 0 {
		return nil, nil, nil
	}
	if len(posts) > limit {
		posts = posts[:limit]
	}
	if len(posts) == 0 {
		return nil, nil, nil
	}
	next := posts[len(posts)-1].ID
	return posts, &next, nil
}
