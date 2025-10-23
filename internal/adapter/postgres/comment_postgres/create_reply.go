package comment_postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"posts_comments_service/internal/domain/models"
	"posts_comments_service/internal/pkg"
)

// CreateReply создает родительский комментарий к посту в базе данных POSTGRES
func (r *CommentRepo) CreateReplyComment(ctx context.Context, postID models.PostID, userID models.UserID, parentCommentID models.CommentID, text string) (*models.Comment, error) {

	var comment models.Comment
	comment.PostID = postID
	comment.UserID = userID
	comment.Text = text
	comment.ParentCommentID = &parentCommentID

	err := pkg.WithTx(ctx, r.cluster.Pool, func(tx pgx.Tx) error {
		row := tx.QueryRow(ctx, `
			INSERT INTO comments (post_id, user_id, parent_comment_id,first_comment_id, text, created_at)
			SELECT $1,$2,$3, c.first_comment_id, $4, NOW()
			FROM comments c
			WHERE c.id = $3 AND c.post_id =$1
			RETURNING id,first_comment_id, created_at`,
			postID, userID, parentCommentID, text)

		if err := row.Scan(&comment.ID, &comment.FirstCommentID, &comment.CreatedAt); err != nil {
			return fmt.Errorf("CreateReplyComment: %w", err)
		}
		return nil

	})
	if err != nil {
		return nil, err
	}
	return &comment, nil

}
