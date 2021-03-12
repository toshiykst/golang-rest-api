package repository

import "gorm.io/gorm"

// DBHandler is interface to handle db
type DBHandler interface {
	Find(interface{}, ...interface{}) *gorm.DB
	Create(interface{}) *gorm.DB
	Update(interface{}) *gorm.DB
	Delete(interface{}) *gorm.DB
}
