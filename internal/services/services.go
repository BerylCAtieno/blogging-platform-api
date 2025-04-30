package services

type Post struct {
	ID       int      `json:"id"`
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Category string   `json:"category"`
	Tags     []string `json:"tags"`
}

type PostStore struct {
	Posts  []Post
	NextID int
}

// get all posts

func (ps *PostStore) GetAllPosts() []Post {
	return ps.Posts
}

// get post by id

// create post
// update post
// patch post
// delete post
// search posts
