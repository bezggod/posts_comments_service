package comment_postgres

import (
	"context"
	"fmt"
	"posts_commets_service/internal/domain/models"
)

func (r *CommentRepo) CreateRoot(ctx context.Context, postID models.PostID, userID models.UserID, text string) (*models.Comment, error) {
	var comment models.Comment
	comment.PostID = postID
	comment.UserID = userID
	comment.Text = text

	row := r.cluster.Conn.QueryRow(ctx, `
	INSERT INTO comments (post_id, user_id, text, created_at)
	VALUES ($1,$2,$3, NOW())
	RETURNING id,created_at
	`, postID, userID, text)

	if err := row.Scan(&comment.ID, &comment.CreatedAt); err != nil {
		return nil, fmt.Errorf("failed to insert comment: %w", err)
	}

	if _, err := r.cluster.Conn.Exec(ctx, `UPDATE comments SET first_comment_id = id WHERE id = $1`, comment.ID); err != nil {
		return nil, err
	}

	comment.FirstCommentID = comment.ID
	return &comment, nil
}
