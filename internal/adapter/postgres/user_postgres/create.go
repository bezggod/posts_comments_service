package user_postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	"posts_comments_service/internal/domain/models"
	"posts_comments_service/internal/pkg"
)

// Create создает нового пользователя в базе данных POSTGRES
func (r *UserRepo) Create(ctx context.Context, name string) (*models.User, error) {
	var user models.User
	user.Name = name

	err := pkg.WithTx(ctx, r.cluster.Pool, func(tx pgx.Tx) error {
		row := r.cluster.Pool.QueryRow(ctx, `INSERT INTO users (name) VALUES ($1) RETURNING id`, name)

		if err := row.Scan(&user.ID); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &user, nil

}
