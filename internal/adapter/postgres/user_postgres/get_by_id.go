package user_postgres

import (
	"context"
	"posts_comments_service/internal/domain/models"
)

// GetByID возвращает пользователя по ID из базы данных POSTGRES
func (r *UserRepo) GetByID(ctx context.Context, id models.UserID) (*models.User, error) {
	row := r.cluster.Pool.QueryRow(ctx, "SELECT id,name FROM users WHERE id = $1", id)

	var user models.User
	if err := row.Scan(&user.ID, &user.Name); err != nil {
		return nil, err
	}
	return &user, nil
}
