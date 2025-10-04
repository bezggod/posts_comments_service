package post_usecase

import "posts_commets_service/internal/domain/models"

func (u *PostUseCase) CommentBlock(id models.PostID, off bool) error {
	return u.posts.CommentBlock(id, off)
}
