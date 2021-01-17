package domain

import (
	"gorm.io/gorm"
)

// Post is a structure of the post.
type Post struct {
	gorm.Model
	Title   string
	Content string
}

// Posts is a set of posts.
type Posts []Post
