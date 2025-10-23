package post_usecase

import "posts_comments_service/internal/domain/interfaces"

type PostUseCase struct {
	posts interfaces.PostRepo
}

func NewPostUseCase(posts interfaces.PostRepo) *PostUseCase {
	return &PostUseCase{posts: posts}
}
