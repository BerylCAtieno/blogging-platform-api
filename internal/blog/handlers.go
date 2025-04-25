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

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	// Decode new data from request body
	var updatedPost Post
	err = json.NewDecoder(r.Body).Decode(&updatedPost)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Find and update post in slice
	for i, post := range posts {
		if post.ID == id {
			updatedPost.ID = id    // Keep same ID
			posts[i] = updatedPost // Update it
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedPost)
			return
		}
	}

	// Post not found
	http.Error(w, "Post not found", http.StatusNotFound)
}

// Patch Post

func PatchPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get ID from query param
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	// Decode the request body into a map (to allow partial updates)
	var updates map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Find and patch the post
	for i := range posts {
		if posts[i].ID == id {
			// Update only the fields provided
			if title, ok := updates["title"].(string); ok {
				posts[i].Title = title
			}
			if content, ok := updates["content"].(string); ok {
				posts[i].Content = content
			}
			if category, ok := updates["category"].(string); ok {
				posts[i].Category = category
			}
			if tags, ok := updates["tags"].([]interface{}); ok {
				var tagStrings []string
				for _, tag := range tags {
					if t, ok := tag.(string); ok {
						tagStrings = append(tagStrings, t)
					}
				}
				posts[i].Tags = tagStrings
			}

			// Respond with the updated post
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(posts[i])
			return
		}
	}

	http.Error(w, "Post not found", http.StatusNotFound)
}

// api handler to delete an existing post

func DeletePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse id from query: /posts?id=2
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	// Find and delete post
	for i, post := range posts {
		if post.ID == id {
			// Delete by slicing out the element
			posts = append(posts[:i], posts[i+1:]...)

			w.WriteHeader(http.StatusNoContent) // 204 No Content
			return
		}
	}

	http.Error(w, "Post not found", http.StatusNotFound)
}

// filter blog posts by a search item

// filter by tag (including multiple tag combos)
