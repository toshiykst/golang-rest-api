package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/toshiykst/golang-rest-api/app/domain/model"
	"github.com/toshiykst/golang-rest-api/app/usecase/mock_usecase"
)

func TestPostServiceHandler_GetPosts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostUsecase := mock_usecase.NewMockPostUsecase(ctrl)

	want := model.Posts{
		{ID: 1, Title: "title1", Content: "content1"},
		{ID: 2, Title: "title2", Content: "content2"},
		{ID: 3, Title: "title3", Content: "content3"}}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/posts")

	mockPostUsecase.EXPECT().GetPosts().Return(want, nil)

	handler := &postServiceHandler{pu: mockPostUsecase}
	err := handler.getPosts(c)

	if err != nil {
		t.Fatalf("want no err, but has error %#v", err)
	}
	if rec.Code != http.StatusOK {
		t.Fatalf("want rec.Code is http.StatusOK, but %#v", rec.Code)
	}
}

func TestPostServiceHandler_GetPosts_InternalServerError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostUsecase := mock_usecase.NewMockPostUsecase(ctrl)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/posts")

	mockPostUsecase.EXPECT().GetPosts().Return(model.Posts{}, errors.New("an error occurred"))

	handler := &postServiceHandler{pu: mockPostUsecase}
	err := handler.getPosts(c)

	if err == nil {
		t.Fatal("want err, but has no error")
	}
	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("want rec.Code is http.StatusInternalServerError, but %#v", rec.Code)
	}
}

func TestPostServiceHandler_GetPost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostUsecase := mock_usecase.NewMockPostUsecase(ctrl)

	p := model.Post{
		ID:      1,
		Title:   "title1",
		Content: "content1",
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/posts/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockPostUsecase.EXPECT().GetPost(p.ID).Return(p, nil)

	handler := &postServiceHandler{pu: mockPostUsecase}
	err := handler.getPost(c)

	if err != nil {
		t.Fatalf("want no err, but has error %#v", err)
	}
	if rec.Code != http.StatusOK {
		t.Fatalf("want rec.Code is http.StatusOK, but %#v", rec.Code)
	}
}

func TestPostServiceHandler_GetPost_InternalServerError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostUsecase := mock_usecase.NewMockPostUsecase(ctrl)

	p := model.Post{
		ID:      1,
		Title:   "title1",
		Content: "content1",
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/posts/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockPostUsecase.EXPECT().GetPost(p.ID).Return(model.Post{}, errors.New("an error occurred"))

	handler := &postServiceHandler{pu: mockPostUsecase}
	err := handler.getPost(c)

	if err == nil {
		t.Fatal("want err, but has no error")
	}
	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("want rec.Code is http.StatusInternalServerError, but %#v", rec.Code)
	}
}

func TestPostServiceHandler_CreatePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostUsecase := mock_usecase.NewMockPostUsecase(ctrl)

	p := &model.Post{}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockPostUsecase.EXPECT().CreatePost(p).Return(nil)

	handler := &postServiceHandler{pu: mockPostUsecase}
	err := handler.createPost(c)

	if err != nil {
		t.Fatalf("want no err, but has error %#v", err)
	}
	if rec.Code != http.StatusCreated {
		t.Fatalf("want rec.Code is http.StatusCreated, but %#v", rec.Code)
	}
}

func TestPostServiceHandler_CreatePost_InternalServerError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostUsecase := mock_usecase.NewMockPostUsecase(ctrl)

	p := &model.Post{}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockPostUsecase.EXPECT().CreatePost(p).Return(errors.New("an error occurred"))

	handler := &postServiceHandler{pu: mockPostUsecase}
	err := handler.createPost(c)

	if err == nil {
		t.Fatal("want err, but has no error")
	}
	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("want rec.Code is http.StatusInternalServerError, but %#v", rec.Code)
	}
}
func TestPostServiceHandler_UpdatePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostUsecase := mock_usecase.NewMockPostUsecase(ctrl)

	want := model.Post{
		ID:      1,
		Title:   "updated title",
		Content: "updated content",
	}

	postJSON, _ := json.Marshal(want)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(postJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/posts/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockPostUsecase.EXPECT().UpdatePost(&want).Return(nil)

	handler := &postServiceHandler{pu: mockPostUsecase}
	err := handler.updatePost(c)

	if err != nil {
		t.Fatalf("want no err, but has error %#v", err)
	}
	if rec.Code != http.StatusOK {
		t.Fatalf("want rec.Code is http.StatusOK, but %#v", rec.Code)
	}
}

func TestPostServiceHandler_UpdatePost_InternalServerError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostUsecase := mock_usecase.NewMockPostUsecase(ctrl)

	want := model.Post{
		ID:      1,
		Title:   "updated title",
		Content: "updated content",
	}

	postJSON, _ := json.Marshal(want)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(postJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/posts/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockPostUsecase.EXPECT().UpdatePost(&want).Return(errors.New("an error occurred"))

	handler := &postServiceHandler{pu: mockPostUsecase}
	err := handler.updatePost(c)

	if err == nil {
		t.Fatal("want err, but has no error")
	}
	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("want rec.Code is http.StatusInternalServerError, but %#v", rec.Code)
	}
}
func TestPostServiceHandler_DeletePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostUsecase := mock_usecase.NewMockPostUsecase(ctrl)

	want := model.Post{
		ID: 1,
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/posts/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockPostUsecase.EXPECT().DeletePost(&want).Return(nil)

	handler := &postServiceHandler{pu: mockPostUsecase}
	err := handler.deletePost(c)

	if err != nil {
		t.Fatalf("want no err, but has error %#v", err)
	}
	if rec.Code != http.StatusOK {
		t.Fatalf("want rec.Code is http.StatusOK, but %#v", rec.Code)
	}
}

func TestPostServiceHandler_DeletePost_InternalServerError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostUsecase := mock_usecase.NewMockPostUsecase(ctrl)

	want := model.Post{
		ID: 1,
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/posts/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockPostUsecase.EXPECT().DeletePost(&want).Return(errors.New("an error occurred"))

	handler := &postServiceHandler{pu: mockPostUsecase}
	err := handler.deletePost(c)

	if err == nil {
		t.Fatal("want err, but has no error")
	}
	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("want rec.Code is http.StatusInternalServerError, but %#v", rec.Code)
	}
}
