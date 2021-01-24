package interactor

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/toshiykst/golang-rest-api/app/domain"
)

type mockPostRepository struct{}

func (r *mockPostRepository) FindPost(id int) (domain.Post, error) {
	return domain.Post{ID: 1, Title: "title1", Content: "content1"}, nil
}

func (r *mockPostRepository) CreatePost(p *domain.Post) error {
	return nil
}

func (r *mockPostRepository) UpdatePost(p *domain.Post) error {
	return nil
}

func (r *mockPostRepository) DeletePost(p *domain.Post) error {
	return nil
}

func (r *mockPostRepository) FindPosts() (domain.Posts, error) {
	return domain.Posts{
			{ID: 1, Title: "title1", Content: "content1"},
			{ID: 2, Title: "title2", Content: "content2"},
			{ID: 3, Title: "title3", Content: "content3"}},
		nil
}

func TestPostInteractor_GetPost(t *testing.T) {
	i := &PostInteractor{
		PostRepository: &mockPostRepository{},
	}
	id := 1
	p, err := i.GetPost(id)

	if err != nil {
		t.Fatalf("want no err, but has error %#v", err)
	}

	want := domain.Post{ID: 1, Title: "title1", Content: "content1"}

	if diff := cmp.Diff(p, want); diff != "" {
		t.Errorf("TestPostInteractor_GetPost differs: (-got +want)\n%s", diff)
	}
}

func TestPostInteractor_GetPosts(t *testing.T) {
	i := &PostInteractor{
		PostRepository: &mockPostRepository{},
	}

	p, err := i.GetPosts()

	if err != nil {
		t.Fatalf("want no err, but has error %#v", err)
	}

	want := domain.Posts{
		{ID: 1, Title: "title1", Content: "content1"},
		{ID: 2, Title: "title2", Content: "content2"},
		{ID: 3, Title: "title3", Content: "content3"}}

	if diff := cmp.Diff(p, want); diff != "" {
		t.Errorf("TestPostInteractor_GetPost differs: (-got +want)\n%s", diff)
	}
}

func TestPostInteractor_CreatePost(t *testing.T) {
	i := &PostInteractor{
		PostRepository: &mockPostRepository{},
	}

	err := i.CreatePost(&domain.Post{Title: "title", Content: "content"})

	if err != nil {
		t.Fatalf("want no err, but has error %#v", err)
	}
}

func TestPostInteractor_UpdatePost(t *testing.T) {
	i := &PostInteractor{
		PostRepository: &mockPostRepository{},
	}

	err := i.UpdatePost(&domain.Post{Title: "updated title", Content: "updated content"})

	if err != nil {
		t.Fatalf("want no err, but has error %#v", err)
	}
}

func TestPostInteractor_DeletePost(t *testing.T) {
	i := &PostInteractor{
		PostRepository: &mockPostRepository{},
	}

	err := i.DeletePost(&domain.Post{Title: "title", Content: "content"})

	if err != nil {
		t.Fatalf("want no err, but has error %#v", err)
	}
}
