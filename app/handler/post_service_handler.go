package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/toshiykst/golang-rest-api/app/domain/model"
	"github.com/toshiykst/golang-rest-api/app/domain/repository"
	"github.com/toshiykst/golang-rest-api/app/infrastructure/datastore"

	"github.com/toshiykst/golang-rest-api/app/usecase"
)

type postServiceHandler struct {
	pu usecase.PostUsecase
}

// HandlePostService handles API Requests for postservice
func HandlePostService(e *echo.Echo, dbHandler repository.DBHandler) {
	pd := datastore.NewPostDataStore(dbHandler)
	pu := usecase.NewPostUsecase(pd)
	handler := postServiceHandler{pu: pu}

	e.GET("/posts", handler.getPosts)
	e.GET("/posts/:id", handler.getPost)
	e.POST("/posts", handler.createPost)
	e.PUT("/posts/:id", handler.updatePost)
	e.DELETE("/posts/:id", handler.deletePost)
}

func (h *postServiceHandler) getPosts(c echo.Context) error {
	ps, err := h.pu.GetPosts()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}

	return c.JSON(http.StatusOK, ps)
}

func (h *postServiceHandler) getPost(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	p, err := h.pu.GetPost(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}

	return c.JSON(http.StatusOK, p)
}

func (h *postServiceHandler) createPost(c echo.Context) error {
	p := &model.Post{}
	c.Bind(p)

	if err := h.pu.CreatePost(p); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}

	return c.JSON(http.StatusCreated, *p)
}

func (h *postServiceHandler) updatePost(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}

	p := &model.Post{ID: id}
	c.Bind(p)

	if err = h.pu.UpdatePost(p); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}

	return c.JSON(http.StatusOK, *p)
}

func (h *postServiceHandler) deletePost(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}

	p := &model.Post{ID: id}
	c.Bind(p)

	if err = h.pu.DeletePost(p); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}

	return c.JSON(http.StatusOK, *p)
}
