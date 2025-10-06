package comment_postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	"posts_commets_service/internal/domain/models"
)

func (r *CommentRepo) ListRoots(ctx context.Context, postID models.PostID, limit int, lastID *models.CommentID) ([]*models.Comment, *models.CommentID, error) {

	var rows pgx.Rows
	var err error

	if lastID == nil {
		rows, err = r.cluster.Conn.Query(ctx, `SELECT id, post_id,user_id, text, created_at
FROM comments where post_id = $1 AND parent_id IS NULL ORDER BY id DESC LIMIT $2`, *lastID, limit)
	} else {
		rows, err = r.cluster.Conn.Query(ctx, `SELECT id, post_id,user_id, text, created_at
FROM comments where post_id = $1 AND parent_id IS NULL ORDER BY id DESC LIMIT $2`, postID, limit)
	}
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var comments []*models.Comment
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Text, &comment.CreatedAt)
		if err != nil {
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
