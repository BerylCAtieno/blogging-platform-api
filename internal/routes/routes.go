package routes

import (
	"blogging-platform-api/internal/blog"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	postsRouter := router.PathPrefix("/posts").Subrouter()

	postsRouter.HandleFunc("", blog.GetPosts).Methods("GET")
	postsRouter.HandleFunc("", blog.CreatePost).Methods("POST")
	postsRouter.HandleFunc("/{id}", blog.GetPostById).Methods("GET")
	postsRouter.HandleFunc("/{id}", blog.UpdatePost).Methods("PUT")
	postsRouter.HandleFunc("/{id}", blog.PatchPost).Methods("PATCH")
	postsRouter.HandleFunc("/{id}", blog.DeletePost).Methods("DELETE")

	return router
}
