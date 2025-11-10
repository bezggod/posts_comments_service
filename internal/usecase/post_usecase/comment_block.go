package post_usecase

import (
	"context"
	"posts_comments_service/internal/domain/models"
)

// CommentBlock дает возможность заблокировать написание комментария
func (u *PostUseCase) CommentBlock(ctx context.Context, id models.PostID, off bool) error {
	return u.posts.CommentBlock(ctx, id, off)
}
