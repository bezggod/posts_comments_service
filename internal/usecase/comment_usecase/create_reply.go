package comment_usecase

import (
	"context"
	"errors"
	"posts_commets_service/internal/domain/models"
)

func (u *CommentUseCase) CreateReplyComment(ctx context.Context, postID models.PostID, userID models.UserID, parentCommentID models.CommentID, text string) (*models.Comment, error) {
	if text == "" {
		return nil, errors.New("text is empty")
	}
	post, err := u.posts.GetByID(ctx, postID)
	if err != nil {
		return nil, err
	}
	if post.CommentBlock {
		return nil, errors.New("post comment is blocked")
	}
	return u.comments.CreateReplyComment(ctx, postID, userID, parentCommentID, text)
}
