package comment_usecase

import (
	"context"
	"fmt"
	"posts_comments_service/internal/domain/models"
)

// CreateRoot создает корневой комментарий к посту
func (u *CommentUseCase) CreateRoot(ctx context.Context, postID models.PostID, userID models.UserID, text string) (*models.Comment, error) {
	post, err := u.posts.GetByID(ctx, postID)
	if err != nil {
		return nil, err
	}
	if post.CommentBlock {
		return nil, fmt.Errorf("post comment is blocked")
	}
	return u.comments.CreateRoot(ctx, postID, userID, text)
}
