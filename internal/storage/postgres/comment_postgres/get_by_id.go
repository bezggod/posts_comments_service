package comment_postgres

import (
	"context"
	"posts_commets_service/internal/domain/models"
)

func (r *CommentRepo) GetByID(ctx context.Context, id models.CommentID) (*models.Comment, error) {
	row := r.cluster.Conn.QueryRow(ctx, `SELECT id, post_id,user_id,parent_comment_id, first_comment_id, text,created_at FROM comments WHERE id = $1`, id)

	var comment models.Comment
	err := row.Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.ParentCommentID, &comment.FirstCommentID, &comment.Text, &comment.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}
