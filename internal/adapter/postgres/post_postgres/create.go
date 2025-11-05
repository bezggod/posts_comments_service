package post_postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"posts_comments_service/internal/domain/models"
	"posts_comments_service/internal/pkg"
)

// Create создает новый пост в базе данных postgres
func (r *PostRepo) Create(ctx context.Context, post *models.Post) (*models.Post, error) {
	if post == nil {
		return nil, fmt.Errorf("post is nil")
	}
	if post.Title == "" {
		return nil, fmt.Errorf("post title is empty")
	}

	var out models.Post
	out.UserID = post.UserID
	out.Title = post.Title
	out.Body = post.Body
	out.CommentBlock = post.CommentBlock

	err := pkg.WithTx(ctx, r.cluster.Pool, func(tx pgx.Tx) error {
		row := tx.QueryRow(ctx, `
			INSERT INTO posts (user_id, title, body, comment_block, created_at)
			VALUES ($1, $2, $3, $4, NOW()) RETURNING id, created_at`,
			post.UserID, post.Title, post.Body, post.CommentBlock)

		if err := row.Scan(&out.ID, &out.CreatedAt); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &out, nil
}
