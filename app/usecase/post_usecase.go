package usecase

import (
	"github.com/toshiykst/golang-rest-api/app/domain/model"
	"github.com/toshiykst/golang-rest-api/app/domain/repository"
)

// PostUsecase is methods
type PostUsecase interface {
	GetPost(id int) (model.Post, error)
	GetPosts() (model.Posts, error)
	CreatePost(*model.Post) error
	UpdatePost(*model.Post) error
	DeletePost(*model.Post) error
}

type postUsecase struct {
	r repository.PostRepository
}

// NewPostUsecase creates a post usecase instance
func NewPostUsecase(r repository.PostRepository) PostUsecase {
	return &postUsecase{
		r: r,
	}
}

func (pu *postUsecase) GetPosts() (post model.Posts, err error) {
	post, err = pu.r.FindPosts()
	return
}

func (pu *postUsecase) GetPost(id int) (post model.Post, err error) {
	post, err = pu.r.FindPost(id)
	return
}

// CreatePost creates the post.
func (pu *postUsecase) CreatePost(p *model.Post) (err error) {
	err = pu.r.CreatePost(p)
	return
}

// UpdatePost updates the post.
func (pu *postUsecase) UpdatePost(p *model.Post) (err error) {
	err = pu.r.UpdatePost(p)
	return
}

// DeletePost deletes the post.
func (pu *postUsecase) DeletePost(p *model.Post) (err error) {
	err = pu.r.DeletePost(p)
	return
}
