package routes

import (
	"blogging-platform-api/internal/blog"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	v1Router := router.PathPrefix("/api/v1").Subrouter()

	postsRouter := v1Router.PathPrefix("/posts").Subrouter()

	postsRouter.HandleFunc("/", blog.GetPosts).Methods("GET").Name("getposts")
	postsRouter.HandleFunc("/", blog.CreatePost).Methods("POST").Name("createpost")
	postsRouter.HandleFunc("/{id:[0-9]+}", blog.GetPostById).Methods("GET").Name("getpostbyid")
	postsRouter.HandleFunc("/{id:[0-9]+}", blog.UpdatePost).Methods("PUT").Name("updatepost")
	postsRouter.HandleFunc("/{id:[0-9]+}", blog.PatchPost).Methods("PATCH").Name("patchpost")
	postsRouter.HandleFunc("/{id:[0-9]+}", blog.DeletePost).Methods("DELETE").Name("deletepost")
	postsRouter.HandleFunc("/search", blog.SearchPosts).Methods("GET").Name("searchposts")

	return router
}
