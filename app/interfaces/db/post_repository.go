package db

import "github.com/toshiykst/golang-rest-api/app/domain"

// PostRepository is a structure of PostRepository.
type PostRepository struct {
	DB
}

// FindPost returns the post by id.
func (r *PostRepository) FindPost(id int) (Post domain.Post, err error) {
	if err = r.Find(&Post, id).Error; err != nil {
		return
	}
	return
}

// CreatePost creates the post.
func (r *PostRepository) CreatePost(p domain.Post) (Post domain.Post, err error) {
	if err = r.Create(&p).Error; err != nil {
		return
	}
	Post = p
	return
}

// UpdatePost updates the post.
func (r *PostRepository) UpdatePost(p domain.Post) (Post domain.Post, err error) {
	if err = r.Update(&p).Error; err != nil {
		return
	}
	Post = p
	return
}

// DeletePost deletes the post.
func (r *PostRepository) DeletePost(p domain.Post) (err error) {
	if err = r.Delete(&p).Error; err != nil {
		return
	}
	return
}

// FindPosts returns posts.
func (r *PostRepository) FindPosts() (ps domain.Posts, err error) {
	if err = r.Find(&ps).Error; err != nil {
		return
	}
	return
}
