package interractor

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
	Post, err = interactor.PostRepository.Find(id)
	return
}

// Posts returns the Post list.
func (interactor *PostInteractor) Posts() (Posts domain.Posts, err error) {
	Posts, err = interactor.PostRepository.FindAll()
	return
}

// Create creates the Post.
func (interactor *PostInteractor) Create(u domain.Post) (Post domain.Post, err error) {
	Post, err = interactor.PostRepository.Create(u)
	return
}

// Update updates the Post.
func (interactor *PostInteractor) Update(u domain.Post) (Post domain.Post, err error) {
	Post, err = interactor.PostRepository.Update(u)
	return
}

// Delete deletes the Post.
func (interactor *PostInteractor) Delete(Post domain.Post) (err error) {
	err = interactor.PostRepository.Delete(Post)
	return
}
