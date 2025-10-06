package post_postgres

import (
	"posts_commets_service/internal/config"
)

type PostRepo struct {
	cluster *config.Cluster
}

func NewPostRepo(cluster *config.Cluster) *PostRepo {
	return &PostRepo{
		cluster: cluster,
	}
}
