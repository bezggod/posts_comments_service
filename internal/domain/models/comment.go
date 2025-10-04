package models

import "time"

type CommentID int64
type Comment struct {
	ID              CommentID
	PostID          PostID
	UserID          UserID
	Text            string
	ParentCommentID *CommentID
	FirstCommentID  CommentID
	CreatedAt       time.Time
}

func NewComment(postID PostID, userID UserID, text string) *Comment {
	return &Comment{
		PostID:    postID,
		UserID:    userID,
		Text:      text,
		CreatedAt: time.Time{},
	}
}
func NewCommentReply(postID PostID, userID UserID, parentID CommentID, rootID CommentID, text string) *Comment {
	return &Comment{
		PostID:          postID,
		UserID:          userID,
		ParentCommentID: &parentID,
		FirstCommentID:  rootID,
		Text:            text,
		CreatedAt:       time.Now(),
	}
}
