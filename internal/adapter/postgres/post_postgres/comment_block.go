package post_postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"posts_comments_service/internal/domain/models"
	"posts_comments_service/internal/pkg"
)

// CommentBlock запрещает или разрешает написание комментария поста в базе данных POSTGRES
func (r *PostRepo) CommentBlock(ctx context.Context, id models.PostID, off bool) error {

	err := pkg.WithTx(ctx, r.cluster.Pool, func(tx pgx.Tx) error {
		_, err := tx.Exec(ctx, `UPDATE posts SET comment_block = $1 WHERE id = $2`, off, id)
		if err != nil {
			return fmt.Errorf("CommentBlock: %w", err)
		}
		return nil
	})
	return err
}
