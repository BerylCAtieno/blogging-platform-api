# blogging-platform-api
A simple RESTful API with basic CRUD operations for a blogging platform

/blogapi
│
├── main.go                      // Entry point
│
├── /router                      // Setup for Gin or Mux routing
│   └── router.go
│
├── /blog
│   ├── handlers.go              // HTTP handlers for blog posts
│   ├── service.go               // Business logic
│   ├── model.go                 // Data models (Post, Comment, etc.)
│   └── repository.go            // DB access
│
├── /user
│   ├── handlers.go
│   ├── service.go
│   └── model.go
│
├── /middleware
│   └── auth.go                  // e.g., Auth middleware
│
└── /utils
    └── helper.go                // Common utilities

