package user_postgres

import (
	"context"
	"posts_commets_service/internal/domain/models"
)

func (r *UserRepo) Create(ctx context.Context, name string) (*models.User, error) {
	row := r.cluster.Conn.QueryRow(ctx, `INSERT INTO users (name) VALUES ($1) RETURNING id`, name)

	var user models.User
	user.Name = name
	if err := row.Scan(&user.ID); err != nil {
		return nil, err
	}
	return &user, nil
}
