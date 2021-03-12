package dao

import (
	"github.com/toshiykst/golang-rest-api/app/domain/model"
	"github.com/toshiykst/golang-rest-api/app/domain/repository"
)

// PostDao is a structure of PostDao.
type PostDao struct {
	DBHandler repository.DBHandler
}

// PostDao creates an instance of PostDao
func NewPostDao(dbHandler repository.DBHandler) *PostDao {
	return &PostDao{
		DBHandler: dbHandler,
	}
}

// FindPost returns the post by id.
func (d *PostDao) FindPost(id int) (p model.Post, err error) {
	if err = d.DBHandler.Find(&p, id).Error; err != nil {
		return
	}
	return
}

// CreatePost creates the post.
func (d *PostDao) CreatePost(p *model.Post) (err error) {
	if err = d.DBHandler.Create(p).Error; err != nil {
		return
	}
	return
}

// UpdatePost updates the post.
func (d *PostDao) UpdatePost(p *model.Post) (err error) {
	if err = d.DBHandler.Update(p).Error; err != nil {
		return
	}
	return
}

// DeletePost deletes the post.
func (d *PostDao) DeletePost(p *model.Post) (err error) {
	if err = d.DBHandler.Delete(p).Error; err != nil {
		return
	}
	return
}

// FindPosts returns posts.
func (d *PostDao) FindPosts() (ps model.Posts, err error) {
	if err = d.DBHandler.Find(&ps).Error; err != nil {
		return
	}
	return
}
