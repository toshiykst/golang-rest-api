package infrastructure

import (
	"fmt"

	"github.com/toshiykst/golang-rest-api/app/interfaces/repository"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Repository is a structure of the repository.
type Repository struct {
	Conn *gorm.DB
}

// NewRepository creates repository.
func NewRepository() repository.Repository {

	connection := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		"user",
		"password",
		"db",
		"3306",
		"golang_rest_api",
	)
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.Set("gorm:table_options", "ENGINE=InnoDB")

	repository := new(Repository)

	repository.Conn = db

	return repository
}

// Find returns the record by query.
func (handler *Repository) Find(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.Find(out, where...)
}

// Create creates the new record.
func (handler *Repository) Create(value interface{}) *gorm.DB {
	return handler.Conn.Create(value)
}

// Save updates the record.
func (handler *Repository) Save(value interface{}) *gorm.DB {
	return handler.Conn.Save(value)
}

// Delete deletes the record.
func (handler *Repository) Delete(value interface{}) *gorm.DB {
	return handler.Conn.Delete(value)
}
