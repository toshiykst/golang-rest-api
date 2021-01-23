package controllers

import (
	"net/http"
	"strconv"

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
	p := &domain.Post{}
	c.Bind(p)

	if err = controller.Interactor.CreatePost(p); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, *p)
	return
}

// GetPosts returns posts
func (controller *PostController) GetPosts(c Context) (err error) {
	ps, err := controller.Interactor.GetPosts()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, ps)
	return
}

// GetPost returns post
func (controller *PostController) GetPost(c Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	p, err := controller.Interactor.GetPost(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, p)
	return
}

// UpdatePost updates post
func (controller *PostController) UpdatePost(c Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	p := &domain.Post{ID: id}
	c.Bind(p)

	if err = controller.Interactor.UpdatePost(p); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, *p)
	return
}

// DeletePost deletes post
func (controller *PostController) DeletePost(c Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	p := &domain.Post{ID: id}

	if err = controller.Interactor.DeletePost(p); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
	return
}
