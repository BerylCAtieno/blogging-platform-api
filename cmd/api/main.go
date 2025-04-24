package main

import (
	"blogging-platform-api/internal/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.SetupRoutes()
	log.Println("API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
