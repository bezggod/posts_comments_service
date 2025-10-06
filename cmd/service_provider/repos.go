package service_provider

import (
	"context"
	"os"
	"posts_commets_service/internal/domain/interfaces"
	"posts_commets_service/internal/storage/in_memory/comment_storage"
	"posts_commets_service/internal/storage/in_memory/post_storage"
	"posts_commets_service/internal/storage/in_memory/user_storage"
	"posts_commets_service/internal/storage/postgres/comment_postgres"
	"posts_commets_service/internal/storage/postgres/post_postgres"
	"posts_commets_service/internal/storage/postgres/user_postgres"
)

func (s *ServiceProvider) getUserRepo() interfaces.UserRepo {
	if s.userRepo == nil {
		if os.Getenv("STORAGE_MODE") == "postgres" {
			s.userRepo = user_postgres.NewUserRepo(s.getDbCluster(context.Background()))
		} else {
			s.userRepo = user_storage.NewUserRepo()
		}
	}
	return s.userRepo
}

func (s *ServiceProvider) getPostRepo() interfaces.PostRepo {
	if s.postRepo == nil {
		if os.Getenv("STORAGE_MODE") == "postgres" {
			s.postRepo = post_postgres.NewPostRepo(s.getDbCluster(context.Background()))
		} else {
			s.postRepo = post_storage.NewPostRepo()
		}
	}
	return s.postRepo
}

func (s *ServiceProvider) getCommentRepo() interfaces.CommentRepo {
	if s.commentRepo == nil {
		if os.Getenv("STORAGE_MODE") == "postgres" {
			s.commentRepo = comment_postgres.NewCommentRepo(s.getDbCluster(context.Background()))
		} else {
			s.commentRepo = comment_storage.NewCommentRepo()
		}
	}
	return s.commentRepo
}
