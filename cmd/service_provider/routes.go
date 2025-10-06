package service_provider

import "net/http"

func (s *ServiceProvider) GetHTTPServer() *http.ServeMux {
	mux := http.NewServeMux()
	userCtrl := s.GetUserController()
	postCtrl := s.GetPostController()
	commentCtrl := s.GetCommentController()

	mux.HandleFunc("POST /users", userCtrl.Create)
	mux.HandleFunc("GET /users/{id}", userCtrl.GetByID)

	mux.HandleFunc("POST /posts", postCtrl.Create)
	mux.HandleFunc("GET /posts/{id}", postCtrl.GetByID)
	mux.HandleFunc("GET /posts", postCtrl.ListPosts)
	mux.HandleFunc("POST /posts/comment_block", postCtrl.CommentBlock)

	mux.HandleFunc("POST /comments/reply", commentCtrl.CreateReply)
	mux.HandleFunc("POST /comments/root", commentCtrl.CreateRoot)
	mux.HandleFunc("GET /comments/{id}", commentCtrl.GetByID)
	mux.HandleFunc("GET /comments/roots", commentCtrl.ListRoots)
	mux.HandleFunc("GET /comments/threads", commentCtrl.ListThread)

	return mux

}
