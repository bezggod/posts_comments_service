package comment_postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	"posts_comments_service/internal/domain/models"
)

// ListRoots возвращает список корневых комментариев поста из базы данных
func (r *CommentRepo) ListRoots(ctx context.Context, postID models.PostID, limit int, lastID *models.CommentID) ([]*models.Comment, *models.CommentID, error) {

	var rows pgx.Rows
	var err error

	if lastID == nil {
		rows, err = r.cluster.Pool.Query(ctx, `
			SELECT id, post_id,user_id, text, created_at
			FROM comments 
			WHERE post_id = $1 AND parent_id IS NULL 
			ORDER BY id DESC LIMIT $2`, postID, limit)
	} else {
		rows, err = r.cluster.Pool.Query(ctx, `
			SELECT id, post_id,user_id, text, created_at
			FROM comments 
			WHERE post_id = $1 AND parent_id IS NULL 
			ORDER BY id DESC LIMIT $2`, *lastID, limit)
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
			&comment.Text,
			&comment.CreatedAt); err != nil {
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
