package post_postgres

import (
	"context"
	"fmt"
	"posts_commets_service/internal/domain/models"
)

func (r *PostRepo) CommentBlock(ctx context.Context, id models.PostID, off bool) error {

	_, err := r.cluster.Conn.Exec(ctx, `UPDATE posts SET comment_block = $1 WHERE id = $2`, off, id)
	if err != nil {
		return fmt.Errorf("CommentBlock: %w", err)
	}
	return nil
}
