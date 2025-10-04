package comment_usecase

import (
	"fmt"
	"posts_commets_service/internal/domain/models"
)

func (u *CommentUsecase) CreateRoot(postID models.PostID, userID models.UserID, text string) (*models.Comment, error) {
	if text == "" {
		return nil, fmt.Errorf("text is empty")
	}
	post, err := u.posts.GetById(postID)
	if err != nil {
		return nil, err
	}
	if post.CommentBlock {
		return nil, fmt.Errorf("post comment is blocked")
	}
	return u.comments.CreateRoot(postID, userID, text)
}
