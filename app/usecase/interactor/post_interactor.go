package interactor

import (
	"github.com/toshiykst/golang-rest-api/app/domain"
	"github.com/toshiykst/golang-rest-api/app/usecase/repository"
)

// PostInteractor is a structure of PostInteractor.
type PostInteractor struct {
	PostRepository repository.PostRepository
}

// Post returns a Post that specified by id.
func (interactor *PostInteractor) Post(id int) (Post domain.Post, err error) {
	Post, err = interactor.PostRepository.FindPost(id)
	return
}

// Posts returns the Post list.
func (interactor *PostInteractor) Posts() (Posts domain.Posts, err error) {
	Posts, err = interactor.PostRepository.FindPosts()
	return
}

// Create creates the Post.
func (interactor *PostInteractor) Create(p domain.Post) (Post domain.Post, err error) {
	Post, err = interactor.PostRepository.CreatePost(p)
	return
}

// Update updates the Post.
func (interactor *PostInteractor) Update(p domain.Post) (Post domain.Post, err error) {
	Post, err = interactor.PostRepository.UpdatePost(p)
	return
}

// Delete deletes the Post.
func (interactor *PostInteractor) Delete(p domain.Post) (err error) {
	err = interactor.PostRepository.DeletePost(p)
	return
}
