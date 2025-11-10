package user_postgres

import (
	"posts_comments_service/internal/config"
)

type UserRepo struct {
	cluster *config.Cluster
}

func NewUserRepo(cluster *config.Cluster) *UserRepo {
	return &UserRepo{
		cluster: cluster,
	}
}
