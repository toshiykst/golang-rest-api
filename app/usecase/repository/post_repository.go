package repository

import "github.com/toshiykst/golang-rest-api/app/domain"

// PostRepository is the PostRepository interface.
type PostRepository interface {
	FindPost(id int) (domain.Post, error)
	CreatePost(*domain.Post) error
	UpdatePost(*domain.Post) error
	DeletePost(*domain.Post) error
	FindPosts() (domain.Posts, error)
}
