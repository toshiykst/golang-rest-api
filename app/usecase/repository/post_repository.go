package repository

import "github.com/toshiykst/golang-rest-api/app/domain"

// PostRepository is the PostRepository interface.
type PostRepository interface {
	Find(id int) (domain.Post, error)
	Create(domain.Post) (domain.Post, error)
	Update(domain.Post) (domain.Post, error)
	Delete(domain.Post) error
	FindAll() (domain.Posts, error)
}
