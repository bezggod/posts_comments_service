package comment_storage

import (
	"posts_commets_service/internal/domain/models"
	"sort"
)

func (r *CommentRepo) ListRoots(postID models.PostID, limit int, lastID *models.CommentID) ([]*models.Comment, *models.CommentID, error) {

	r.mu.RLock()
	var comments []*models.Comment
	for _, c := range r.byID {
		if c.PostID == postID && c.ParentCommentID == nil {
			if lastID == nil || c.ID < *lastID {
				comments = append(comments, c)
			}
		}
	}
	r.mu.RUnlock()

	sort.Slice(comments, func(i, j int) bool {
		return comments[i].ID < comments[j].ID
	})
	if limit > 0 && len(comments) == limit {
		comments = comments[:limit]
	}
	if len(comments) == 0 {
		return comments, nil, nil
	}
	next := comments[len(comments)-1].ID
	return comments, &next, nil
}
