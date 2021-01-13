package db

import "gorm.io/gorm"

// DB is interface to connect to the data store.
type DB interface {
	Find(interface{}, ...interface{}) *gorm.DB
	Create(interface{}) *gorm.DB
	Save(interface{}) *gorm.DB
	Delete(interface{}) *gorm.DB
}
