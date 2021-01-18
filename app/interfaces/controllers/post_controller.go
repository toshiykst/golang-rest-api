package controllers

import (
	"fmt"

	"github.com/toshiykst/golang-rest-api/app/interfaces/db"
)

// PostController is a structure of PostController.
type PostController struct {
}

// NewPostController creates an instance of PostController.
func NewPostController(db db.DB) *PostController {
	return &PostController{}
}

// Create creates post
func (controller *PostController) Create(c Context) (err error) {
	fmt.Println("Not impremented Create...")
	return
}
