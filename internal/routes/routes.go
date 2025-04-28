package routes

import (
	"blogging-platform-api/internal/blog"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	postsRouter := router.PathPrefix("/posts").Subrouter()

	postsRouter.HandleFunc("", blog.GetPosts).Methods("GET").Name("getposts")
	postsRouter.HandleFunc("", blog.CreatePost).Methods("POST").Name("createpost")
	postsRouter.HandleFunc("/{id}", blog.GetPostById).Methods("GET").Name("getpostbyid")
	postsRouter.HandleFunc("/{id}", blog.UpdatePost).Methods("PUT").Name("updatepost")
	postsRouter.HandleFunc("/{id}", blog.PatchPost).Methods("PATCH").Name("patchpost")
	postsRouter.HandleFunc("/{id}", blog.DeletePost).Methods("DELETE").Name("deletepost")

	return router
}
