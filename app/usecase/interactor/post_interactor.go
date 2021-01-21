package interactor

import (
	"github.com/toshiykst/golang-rest-api/app/domain"
	"github.com/toshiykst/golang-rest-api/app/usecase/repository"
)

// PostInteractor is a structure of PostInteractor.
type PostInteractor struct {
	PostRepository repository.PostRepository
}

// GetPost returns the Post.
func (interactor *PostInteractor) GetPost(id int) (Post domain.Post, err error) {
	Post, err = interactor.PostRepository.FindPost(id)
	return
}

// GetPosts returns the post list.
func (interactor *PostInteractor) GetPosts() (Posts domain.Posts, err error) {
	Posts, err = interactor.PostRepository.FindPosts()
	return
}

// CreatePost creates the post.
func (interactor *PostInteractor) CreatePost(p domain.Post) (Post domain.Post, err error) {
	Post, err = interactor.PostRepository.CreatePost(p)
	return
}

// UpdatePost updates the post.
func (interactor *PostInteractor) UpdatePost(p domain.Post) (Post domain.Post, err error) {
	Post, err = interactor.PostRepository.UpdatePost(p)
	return
}

// DeletePost deletes the post.
func (interactor *PostInteractor) DeletePost(p domain.Post) (err error) {
	err = interactor.PostRepository.DeletePost(p)
	return
}
