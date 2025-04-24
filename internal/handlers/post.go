package handlers

import (
	"blogging-platform-api/internal/models"
	"encoding/json"
	"net/http"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	posts := []models.Post_Payload{
		{Title: "First Post", Content: "Hello world!", Category: "random", Tags: []string{"tech", "technology"}},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
