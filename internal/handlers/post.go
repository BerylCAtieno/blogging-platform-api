package handlers

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Category string   `json:"category"`
	Tags     []string `json:"tags"`
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	posts := []Post{
		{Title: "First Post", Content: "Hello world!", Category: "random", Tags: []string{"tech", "technology"}},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
