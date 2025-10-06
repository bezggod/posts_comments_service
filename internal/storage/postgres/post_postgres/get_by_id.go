package post_postgres

import (
	"context"
	"posts_commets_service/internal/domain/models"
)

func (r *PostRepo) GetByID(ctx context.Context, id models.PostID) (*models.Post, error) {
	row := r.cluster.Conn.QueryRow(ctx, "SELECT id,user_id,title,body,comment_block, created_at FROM posts WHERE id = $1", id)
	var posts models.Post
	if err := row.Scan(&posts.ID, &posts.UserID, &posts.Title, &posts.Body, &posts.CommentBlock, &posts.CreatedAt); err != nil {
		return nil, err
	}
	return &posts, nil
}
