package models

import "time"

type UserID int64
type PostID int64

type Post struct {
	ID           PostID
	UserID       UserID
	Title        string
	Body         string
	CommentBlock bool
	CreatedAt    time.Time
	Text         string
}

func NewPost(userID UserID, title, body string, commentBlock bool, createdAt time.Time) *Post {
	return &Post{
		UserID:       userID,
		Title:        title,
		Body:         body,
		CommentBlock: commentBlock,
		CreatedAt:    time.Time{},
	}
}
