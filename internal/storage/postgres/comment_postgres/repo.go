package comment_postgres

import (
	"fmt"
	"posts_commets_service/internal/config"
)

var NotFound = fmt.Errorf("comment not found")

type CommentRepo struct {
	cluster *config.Cluster
}

func NewCommentRepo(cluster *config.Cluster) *CommentRepo {
	return &CommentRepo{
		cluster: cluster,
	}
}
