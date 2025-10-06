package comment_postgres

import (
	"context"
	"fmt"
	"posts_commets_service/internal/domain/models"
)

func (r *CommentRepo) CreateReplyComment(ctx context.Context, postID models.PostID, userID models.UserID, parentCommentID models.CommentID, text string) (*models.Comment, error) {
	row := r.cluster.Conn.QueryRow(ctx, `
	INSERT INTO comments (post_id, user_id, parent_comment_id,first_comment_id, text, created_at)
	SELECT $1,$2,$3, c.first_comment_id, $4, NOW()
	FROM comments c
	WHERE c.id = $3 AND c.post_id =$1
	RETURNING id,first_comment_id, created_at`,
		postID, userID, parentCommentID, text)

	var comment models.Comment
	comment.PostID = postID
	comment.UserID = userID
	comment.Text = text
	comment.ParentCommentID = &parentCommentID

	if err := row.Scan(&comment.ID, &comment.FirstCommentID, &comment.CreatedAt); err != nil {
		return nil, fmt.Errorf("CreateReplyComment: %w", err)
	}
	return &comment, nil
}
