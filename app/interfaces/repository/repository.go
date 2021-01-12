package repository

import "gorm.io/gorm"

// Repository is methods to connect the data store.
type Repository interface {
	Find(interface{}, ...interface{}) *gorm.DB
	Create(interface{}) *gorm.DB
	Save(interface{}) *gorm.DB
	Delete(interface{}) *gorm.DB
}
