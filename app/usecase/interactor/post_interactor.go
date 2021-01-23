package interactor

import (
	"github.com/toshiykst/golang-rest-api/app/domain"
	"github.com/toshiykst/golang-rest-api/app/usecase/repository"
)

// PostInteractor is a structure of PostInteractor.
type PostInteractor struct {
	PostRepository repository.PostRepository
}

// GetPost returns the post.
func (interactor *PostInteractor) GetPost(id int) (p domain.Post, err error) {
	p, err = interactor.PostRepository.FindPost(id)
	return
}

// GetPosts returns the post list.
func (interactor *PostInteractor) GetPosts() (ps domain.Posts, err error) {
	ps, err = interactor.PostRepository.FindPosts()
	return
}

// CreatePost creates the post.
func (interactor *PostInteractor) CreatePost(p *domain.Post) (err error) {
	err = interactor.PostRepository.CreatePost(p)
	return
}

// UpdatePost updates the post.
func (interactor *PostInteractor) UpdatePost(p *domain.Post) (err error) {
	err = interactor.PostRepository.UpdatePost(p)
	return
}

// DeletePost deletes the post.
func (interactor *PostInteractor) DeletePost(p *domain.Post) (err error) {
	err = interactor.PostRepository.DeletePost(p)
	return
}
