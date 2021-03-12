package usecase

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/toshiykst/golang-rest-api/app/domain/model"
	"github.com/toshiykst/golang-rest-api/app/domain/repository/mock_repository"
)

func TestPostUsecase_GetPost(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockRepository := mock_repository.NewMockPostRepository(ctrl)

	want := model.Post{ID: 1, Title: "title1", Content: "content1"}

	mockRepository.EXPECT().FindPost(want.ID).Return(want, nil)

	i := NewPostUsecase(mockRepository)

	p, err := i.GetPost(want.ID)

	if err != nil {
		t.Fatalf("want no err, but has error %#v", err)
	}

	if diff := cmp.Diff(p, want); diff != "" {
		t.Errorf("TestPostUsecase_GetPost differs: (-got +want)\n%s", diff)
	}
}

func TestPostUsecase_GetPosts(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockRepository := mock_repository.NewMockPostRepository(ctrl)

	want := model.Posts{
		{ID: 1, Title: "title1", Content: "content1"},
		{ID: 2, Title: "title2", Content: "content2"},
		{ID: 3, Title: "title3", Content: "content3"}}

	mockRepository.EXPECT().FindPosts().Return(want, nil)

	i := NewPostUsecase(mockRepository)

	p, err := i.GetPosts()

	if err != nil {
		t.Fatalf("want no err, but has error %#v", err)
	}

	if diff := cmp.Diff(p, want); diff != "" {
		t.Errorf("TestPostUsecase_GetPost differs: (-got +want)\n%s", diff)
	}
}

func TestPostUsecase_CreatePost(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockRepository := mock_repository.NewMockPostRepository(ctrl)

	want := model.Post{Title: "title1", Content: "content1"}

	mockRepository.EXPECT().CreatePost(&want).Return(nil)

	i := NewPostUsecase(mockRepository)

	err := i.CreatePost(&want)

	if err != nil {
		t.Fatalf("want no err, but has error %#v", err)
	}
}

func TestPostUsecase_UpdatePost(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockRepository := mock_repository.NewMockPostRepository(ctrl)

	want := model.Post{ID: 1, Title: "updated title", Content: "updated content"}

	mockRepository.EXPECT().UpdatePost(&want).Return(nil)

	i := NewPostUsecase(mockRepository)

	err := i.UpdatePost(&want)

	if err != nil {
		t.Fatalf("want no err, but has error %#v", err)
	}
}

func TestPostUsecase_DeletePost(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockRepository := mock_repository.NewMockPostRepository(ctrl)

	want := model.Post{ID: 1, Title: "title", Content: "content"}

	mockRepository.EXPECT().DeletePost(&want).Return(nil)

	i := NewPostUsecase(mockRepository)

	err := i.DeletePost(&want)

	if err != nil {
		t.Fatalf("want no err, but has error %#v", err)
	}
}
