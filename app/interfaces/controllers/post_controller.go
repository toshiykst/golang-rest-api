package controllers

import (
	"fmt"

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

// Create creates post
func (controller *PostController) Create(c Context) (err error) {
	fmt.Println("Not impremented Create...")
	return
}
