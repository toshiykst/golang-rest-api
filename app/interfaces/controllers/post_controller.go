package controllers

import (
	"github.com/toshiykst/golang-rest-api/app/domain"
	"github.com/toshiykst/golang-rest-api/app/interfaces/db"
	"github.com/toshiykst/golang-rest-api/app/usecase/interactor"
)

// PostController is a structure of PostController.
type PostController struct {
	Interactor interactor.PostInteractor
}

// NewPostController creates an instance of PostController.
func NewPostController(d db.DB) *PostController {
	return &PostController{
		Interactor: interactor.PostInteractor{
			PostRepository: &db.PostRepository{
				DB: d,
			},
		},
	}
}

// CreatePost creates post
func (controller *PostController) CreatePost(c Context) (err error) {
	p := domain.Post{}
	c.Bind(&p)

	post, err := controller.Interactor.CreatePost(p)

	if err != nil {
		c.JSON(500, err.Error())
	}

	c.JSON(201, post)
	return
}
