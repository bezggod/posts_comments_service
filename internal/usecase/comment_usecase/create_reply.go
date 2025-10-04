package comment_usecase

import (
	"errors"
	"posts_commets_service/internal/domain/models"
)

func (u *CommentUsecase) CreateReplyComment(postID models.PostID, userID models.UserID, parentCommentID models.CommentID, text string) (*models.Comment, error) {
	if text == "" {
		return nil, errors.New("text is empty")
	}
	post, err := u.posts.GetById(postID)
	if err != nil {
		return nil, err
	}
	if post.CommentBlock {
		return nil, errors.New("post comment is blocked")
	}
	return u.comments.CreateReplyComment(postID, userID, parentCommentID, text)
}
