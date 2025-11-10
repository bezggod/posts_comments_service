package comment_postgres

import (
	"context"
	"fmt"
	"posts_comments_service/internal/domain/models"
	"posts_comments_service/internal/pkg"

	"github.com/jackc/pgx/v5"
)

// CreateRoot создает ответа на комментарий в базе данных POSTGRES
func (r *CommentRepo) CreateRoot(ctx context.Context, postID models.PostID, userID models.UserID, text string) (*models.Comment, error) {
	var comment models.Comment
	comment.PostID = postID
	comment.UserID = userID
	comment.Text = text

	err := pkg.WithTx(ctx, r.cluster.Pool, func(tx pgx.Tx) error {
		// Вставляем комментарий
		row := tx.QueryRow(ctx, `
			INSERT INTO comments (post_id, user_id, text, created_at)
			VALUES ($1,$2,$3, NOW())
			RETURNING id,created_at
		`, postID, userID, text)

		if err := row.Scan(&comment.ID, &comment.CreatedAt); err != nil {
			return fmt.Errorf("failed to insert comment: %w", err)
		}

		// Обновляем first_comment_id
		if _, err := tx.Exec(ctx, `
			UPDATE comments 
			SET first_comment_id = id 
			WHERE id = $1`, comment.ID); err != nil {
			return fmt.Errorf("failed to update first_comment_id: %w", err)
		}

		comment.FirstCommentID = comment.ID
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &comment, nil
}
