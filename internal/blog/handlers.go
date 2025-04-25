package blog

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Post struct {
	ID       int      `json:"id"`
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Category string   `json:"category"`
	Tags     []string `json:"tags"`
}

// Temporary

var posts = []Post{} // dummy in-memory post store
var nextID = 1       // simple auto-increment ID

// Get all blog posts

func GetPosts(w http.ResponseWriter, r *http.Request) {
	posts := []Post{
		{ID: 1, Title: "First Post", Content: "Hello world!", Category: "random", Tags: []string{"tech", "technology"}},
		{ID: 2, Title: "Second Post", Content: "Hello world Again!", Category: "technology", Tags: []string{"new", "science"}},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

// Get post by id

func GetPostById(w http.ResponseWriter, r *http.Request) {
	posts := []Post{
		{ID: 1, Title: "First Post", Content: "Hello world!", Category: "random", Tags: []string{"tech", "technology"}},
		{ID: 2, Title: "Second Post", Content: "Hello world Again!", Category: "technology", Tags: []string{"new", "science"}},
	}

	// Extract "id" from query parameter: /post?id=1
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Enter a valid id", http.StatusBadRequest)
		return
	}

	// Find post by id
	for _, post := range posts {
		if post.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(post)
			return
		}
	}

	// If not found
	http.Error(w, "Post not found", http.StatusNotFound)
}

// handler to create new blog post

func CreatePost(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method must be POST", http.StatusMethodNotAllowed)
		return
	}

	var newPost Post
	err := json.NewDecoder(r.Body).Decode(&newPost)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newPost.ID = nextID
	nextID++
	posts = append(posts, newPost)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPost)
}

// api handler to update an existing post

// api handler to delete an existing post

// filter blog posts by a search item

// filter by tag (including multiple tag combos)
