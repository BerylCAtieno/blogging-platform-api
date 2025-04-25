package routes

import (
	"blogging-platform-api/internal/blog"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/posts", blog.GetPosts).Methods("GET")
	router.HandleFunc("/post", blog.GetPostById).Methods("GET")
	router.HandleFunc("/posts", blog.CreatePost).Methods("POST")

	return router
}
