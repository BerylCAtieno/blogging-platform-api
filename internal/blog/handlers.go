package blog

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Post struct {
	ID       int      `json:"id"`
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Category string   `json:"category"`
	Tags     []string `json:"tags"`
}

// TODO: (Delete this after implementing db) Temporary

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

	// Extract "id" from path parameter: /posts/{id}

	vars := mux.Vars(r)
	idStr, ok := vars["id"]

	if !ok {
		http.Error(w, "Missing id path", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Enter a valid id", http.StatusBadRequest)
		return
	}

	for _, post := range posts {
		if post.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(post)
			return
		}
	}

	http.Error(w, "Post not found", http.StatusNotFound)
}

// create new blog post

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

// update an existing post

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)

	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "Missing id path", http.StatusBadRequest)
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

	vars := mux.Vars(r)

	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "Missing id path", http.StatusBadRequest)
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
	vars := mux.Vars(r)

	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "Missing id path", http.StatusBadRequest)
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

// TODO: search api not working properly

// filter blog posts by a search item and/or tags

func SearchPosts(w http.ResponseWriter, r *http.Request) {
	// Extract search parameters
	searchTerm := r.URL.Query().Get("q")
	tags := r.URL.Query()["tags"]

	// Filter posts based on search term and tags
	filteredPosts := filterPosts(posts, searchTerm, tags)

	// Handle the case where no posts are found
	if len(filteredPosts) == 0 {
		http.Error(w, "No posts found matching the search criteria", http.StatusNotFound)
		return
	}

	// Otherwise, return the filtered posts
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // Set 200 OK status
	json.NewEncoder(w).Encode(filteredPosts)
}

func filterPosts(posts []Post, searchTerm string, tags []string) []Post {
	var result []Post

	for _, post := range posts {
		matchSearchTerm := searchTerm == "" || strings.Contains(post.Title, searchTerm) || strings.Contains(post.Content, searchTerm)
		matchTags := len(tags) == 0 || containsAny(post.Tags, tags)

		if matchSearchTerm && matchTags {
			result = append(result, post)
		}
	}

	return result
}

func containsAny(tags []string, searchTags []string) bool {
	for _, tag := range searchTags {
		for _, postTag := range tags {
			if postTag == tag {
				return true
			}
		}
	}
	return false
}
