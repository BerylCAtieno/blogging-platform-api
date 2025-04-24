package routes

import (
	"blogging-platform-api/internal/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/posts", handlers.GetPosts).Methods("GET")
	return router
}
