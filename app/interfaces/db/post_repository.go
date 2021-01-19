package db

import "github.com/toshiykst/golang-rest-api/app/domain"

// PostRepository is a structure of PostRepository.
type PostRepository struct {
	db DB
}

// Find returns the post by id.
func (r *PostRepository) Find(id int) (Post domain.Post, err error) {
	if err = r.db.Find(&Post, id).Error; err != nil {
		return
	}
	return
}

// Create creates the post.
func (r *PostRepository) Create(p domain.Post) (Post domain.Post, err error) {
	if err = r.db.Create(&p).Error; err != nil {
		return
	}
	Post = p
	return
}

// Update updates the post.
func (r *PostRepository) Update(p domain.Post) (Post domain.Post, err error) {
	if err = r.db.Update(&p).Error; err != nil {
		return
	}
	Post = p
	return
}

// Delete deletes the post.
func (r *PostRepository) Delete(p domain.Post) (err error) {
	if err = r.db.Delete(&p).Error; err != nil {
		return
	}
	return
}

// FindAll returns posts.
func (r *PostRepository) FindAll() (ps domain.Posts, err error) {
	if err = r.db.Find(&ps).Error; err != nil {
		return
	}
	return
}
