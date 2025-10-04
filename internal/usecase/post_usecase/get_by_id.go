package post_usecase

import "posts_commets_service/internal/domain/models"

func (u *PostUseCase) GetByID(id models.PostID) (*models.Post, error) {
	return u.posts.GetById(id)
}
