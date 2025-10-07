package service_provider

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"posts_commets_service/internal/graphql"
	generated "posts_commets_service/internal/graphql"

	gqlhandler "github.com/99designs/gqlgen/graphql/handler"
)

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

	mux.Handle("/graphql", s.getGraphQLHandler())
	mux.Handle("/", playground.Handler("GraphQL playground", "/graphql"))

	return mux
}

func (s *ServiceProvider) getGraphQLHandler() http.Handler {
	res := &graphql.Resolver{
		UserUC:    s.GetUserUseCase(),
		PostUC:    s.GetPostUseCase(),
		CommentUC: s.GetCommentUseCase(),
	}

	srv := gqlhandler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: res},
		),
	)

	return srv
	return mux

}
