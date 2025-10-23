package post_postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	"posts_comments_service/internal/domain/models"
	"posts_comments_service/internal/pkg"
)

// Create создает новый пост в базе данных postgres
func (r *PostRepo) Create(ctx context.Context, userID models.UserID, title, body string, commentBlock bool) (*models.Post, error) {
	var post models.Post
	post.UserID = userID
	post.Title = title
	post.Body = body
	post.CommentBlock = commentBlock

	err := pkg.WithTx(ctx, r.cluster.Pool, func(tx pgx.Tx) error {
		row := r.cluster.Pool.QueryRow(ctx, `INSERT INTO posts (user_id, title, body, comment_block, created_at)
	values ($1, $2, $3, $4, NOW()) RETURNING id, created_at`, userID, title, body, commentBlock)

		if err := row.Scan(&post.ID, &post.CreatedAt); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &post, nil
}
