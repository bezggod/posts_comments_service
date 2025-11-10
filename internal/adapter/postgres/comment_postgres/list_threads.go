package comment_postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	"posts_comments_service/internal/domain/models"
)

// ListThreads возвращает дерево ответов на комментарий из базы данных Postgres
func (r *CommentRepo) ListThreads(ctx context.Context, postID models.PostID, firstCommentID models.CommentID, limit int, lastID *models.CommentID) ([]*models.Comment, *models.CommentID, error) {

	var rows pgx.Rows
	var err error

	if lastID == nil {
		rows, err = r.cluster.Pool.Query(ctx, `
			SELECT id, post_id,user_id, text, created_at
			FROM comments 
			WHERE post_id = $1 AND first_comment_id=$2 
			ORDER BY id DESC LIMIT $3`, postID, firstCommentID, limit)
	} else {
		rows, err = r.cluster.Pool.Query(ctx, `
			SELECT id, post_id,user_id, text, created_at
			FROM comments 
			WHERE post_id = $1 AND first_comment_id=$2 
			ORDER BY id DESC LIMIT $3`, postID, *lastID, limit)
	}
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var comments []*models.Comment
	for rows.Next() {
		var comment models.Comment
		if err := rows.Scan(
			&comment.ID,
			&comment.PostID,
			&comment.UserID,
			&comment.ParentCommentID,
			&comment.Text,
			&comment.CreatedAt,
		); err != nil {
			return nil, nil, err
		}
		comments = append(comments, &comment)
	}
	if len(comments) == 0 {
		return comments, nil, nil
	}
	next := comments[len(comments)-1].ID
	return comments, &next, nil

}
