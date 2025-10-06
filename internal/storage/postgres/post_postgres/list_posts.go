package post_postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	"posts_commets_service/internal/domain/models"
)

func (r *PostRepo) ListPosts(ctx context.Context, limit int, lastID *models.PostID) ([]*models.Post, *models.PostID, error) {
	var rows pgx.Rows
	var err error

	if lastID != nil {
		rows, err = r.cluster.Conn.Query(ctx, `select id, user_id, title,body,comment_block, created_at from posts
		FROM posts WHERE id < $1 ORDER BY id DESC LIMIT $2 OFFSET $3`, *lastID, limit)
	} else {
		rows, err = r.cluster.Conn.Query(ctx, `select id, user_id, title,body,comment_block, created_at from posts
		FROM posts ORDER BY id DESC LIMIT $1`, limit)
	}
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var posts []*models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Body, &post.CommentBlock, &post.CreatedAt); err != nil {
			return nil, nil, err
		}
		posts = append(posts, &post)
	}

	if len(posts) == 0 {
		return posts, nil, nil
	}
	next := posts[len(posts)-1].ID
	return posts, &next, nil
}
