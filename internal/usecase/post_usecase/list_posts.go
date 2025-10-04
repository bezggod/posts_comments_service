package post_usecase

import "posts_commets_service/internal/domain/models"

func (u *PostUseCase) ListPosts(limit int, lastID *models.PostID) ([]*models.Post, *models.PostID, error) {
	if limit < 0 {
		limit = 0
	}
	return u.posts.ListPosts(limit, lastID)
}
