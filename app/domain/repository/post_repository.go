package repository

import "github.com/toshiykst/golang-rest-api/app/domain/model"

// PostRepository is interface to get posts from datastore.
type PostRepository interface {
	FindPosts() (model.Posts, error)
	FindPost(id int) (model.Post, error)
	CreatePost(*model.Post) error
	UpdatePost(*model.Post) error
	DeletePost(*model.Post) error
}
