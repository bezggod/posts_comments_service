package main

import (
	"fmt"
	"posts_commets_service/internal/storage/in_memory/comment_storage"
	"posts_commets_service/internal/storage/in_memory/post_storage"
	"posts_commets_service/internal/storage/in_memory/user_storage"
	"posts_commets_service/internal/usecase/comment_usecase"
	"posts_commets_service/internal/usecase/post_usecase"
	"posts_commets_service/internal/usecase/user_usecase"
)

func main() {
	postRepo := post_storage.NewPostRepo()
	userRepo := user_storage.NewUserRepo()
	commentRepo := comment_storage.NewCommentRepo()

	postUC := post_usecase.NewPostUseCase(postRepo)
	userUC := user_usecase.NewUserUseCase(userRepo)
	commentUC := comment_usecase.NewCommentUseCase(postRepo, commentRepo)

	user, _ := userUC.Create("Danila")
	fmt.Println(user)

	post, _ := postUC.Create(user.ID, "FIrst post", "yayaya", false)
	fmt.Println(post)

	root, _ := commentUC.CreateRoot(post.ID, user.ID, "ayyayay")
	fmt.Println(root)

	reply, _ := commentUC.CreateReplyComment(post.ID, user.ID, root.ID, "papapap")
	fmt.Println(reply)

	roots, nextRoodID, err := commentUC.ListRoots(post.ID, 100, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(roots)
	fmt.Println(nextRoodID)

	thread, nextTreadID, err := commentUC.ListThread(post.ID, root.ID, 100, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(thread)
	fmt.Println(nextTreadID)
}
