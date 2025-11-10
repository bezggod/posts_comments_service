package comment_usecase

import (
	"context"
	"errors"
	"posts_comments_service/internal/domain/models"
)

// CreateReplyComment создает ответ на комментарий
func (u *CommentUseCase) CreateReplyComment(ctx context.Context, postID models.PostID, userID models.UserID, parentCommentID models.CommentID, text string) (*models.Comment, error) {
	post, err := u.posts.GetByID(ctx, postID)
	if err != nil {
		return nil, err
	}
	if post.CommentBlock {
		return nil, errors.New("post comment is blocked")
	}
	return u.comments.CreateReplyComment(ctx, postID, userID, parentCommentID, text)
}
