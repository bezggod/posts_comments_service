package user_postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"posts_comments_service/internal/domain/models"
	"posts_comments_service/internal/pkg"
)

// Create создает нового пользователя в базе данных POSTGRES
func (r *UserRepo) Create(ctx context.Context, user *models.User) (*models.User, error) {
	if user == nil {
		return nil, fmt.Errorf("user is nil")
	}

	if user.Name == "" {
		return nil, fmt.Errorf("name is empty")
	}

	var out models.User
	out.Name = user.Name

	err := pkg.WithTx(ctx, r.cluster.Pool, func(tx pgx.Tx) error {
		row := tx.QueryRow(ctx, `INSERT INTO users (name) VALUES ($1) RETURNING id`, user.Name)

		if err := row.Scan(&out.ID); err != nil {
			return fmt.Errorf("insert uesr: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &out, nil

}
