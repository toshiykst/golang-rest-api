package datastore

import (
	"github.com/toshiykst/golang-rest-api/app/domain/model"
	"github.com/toshiykst/golang-rest-api/app/domain/repository"
)

// PostDataStore is a structure of PostDataStore.
type PostDataStore struct {
	DBHandler repository.DBHandler
}

// PostDataStore creates an instance of PostDataStore
func NewPostDataStore(dbHandler repository.DBHandler) *PostDataStore {
	return &PostDataStore{
		DBHandler: dbHandler,
	}
}

// FindPost returns the post by id.
func (d *PostDataStore) FindPost(id int) (p model.Post, err error) {
	if err = d.DBHandler.Find(&p, id).Error; err != nil {
		return
	}
	return
}

// CreatePost creates the post.
func (d *PostDataStore) CreatePost(p *model.Post) (err error) {
	if err = d.DBHandler.Create(p).Error; err != nil {
		return
	}
	return
}

// UpdatePost updates the post.
func (d *PostDataStore) UpdatePost(p *model.Post) (err error) {
	if err = d.DBHandler.Update(p).Error; err != nil {
		return
	}
	return
}

// DeletePost deletes the post.
func (d *PostDataStore) DeletePost(p *model.Post) (err error) {
	if err = d.DBHandler.Delete(p).Error; err != nil {
		return
	}
	return
}

// FindPosts returns posts.
func (d *PostDataStore) FindPosts() (ps model.Posts, err error) {
	if err = d.DBHandler.Find(&ps).Error; err != nil {
		return
	}
	return
}
