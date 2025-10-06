package post_usecase

import (
	"context"
	"posts_commets_service/internal/domain/models"
)

func (u *PostUseCase) CommentBlock(ctx context.Context, id models.PostID, off bool) error {
	return u.posts.CommentBlock(ctx, id, off)
}
