package comment_usecase

import (
	"context"
	"fmt"
	"posts_commets_service/internal/domain/models"
)

func (u *CommentUseCase) CreateRoot(ctx context.Context, postID models.PostID, userID models.UserID, text string) (*models.Comment, error) {
	if text == "" {
		return nil, fmt.Errorf("text is empty")
	}
	post, err := u.posts.GetByID(ctx, postID)
	if err != nil {
		return nil, err
	}
	if post.CommentBlock {
		return nil, fmt.Errorf("post comment is blocked")
	}
	return u.comments.CreateRoot(ctx, postID, userID, text)
}
