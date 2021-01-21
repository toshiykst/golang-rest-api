package domain

import "time"

// Post is a structure of the post.
type Post struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// Posts is a set of posts.
type Posts []Post
