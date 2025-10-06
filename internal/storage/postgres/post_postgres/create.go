package post_postgres

import (
	"context"
	"posts_commets_service/internal/domain/models"
)

func (r *PostRepo) Create(ctx context.Context, userID models.UserID, title, body string, commentBlock bool) (*models.Post, error) {
	row := r.cluster.Conn.QueryRow(ctx, `INSERT INTO posts (user_id, title, body, comment_block, created_at)
	values ($1, $2, $3, $4, NOW()) RETURNING id, created_at`, userID, title, body, commentBlock)

	var post models.Post
	post.UserID = userID
	post.Title = title
	post.Body = body
	post.CommentBlock = commentBlock

	if err := row.Scan(&post.ID, &post.CreatedAt); err != nil {
		return nil, err
	}
	return &post, nil
}
