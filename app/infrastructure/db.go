package infrastructure

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/toshiykst/golang-rest-api/app/domain/repository"
)

// DB is a structure of the DB.
type DB struct {
	Conn *gorm.DB
}

// NewDB creates an instance of DB.
func NewDB() repository.DBHandler {

	connection := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		"db",
		"3306",
		os.Getenv("MYSQL_DATABASE"),
	)
	conn, err := gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	conn.Set("gorm:table_options", "ENGINE=InnoDB")

	db := new(DB)
	db.Conn = conn

	return db
}

// Find returns the record by query.
func (db *DB) Find(out interface{}, where ...interface{}) *gorm.DB {
	return db.Conn.Find(out, where...)
}

// Create creates the new record.
func (db *DB) Create(value interface{}) *gorm.DB {
	return db.Conn.Create(value)
}

// Update updates the record.
func (db *DB) Update(value interface{}) *gorm.DB {
	return db.Conn.Save(value)
}

// Delete deletes the record.
func (db *DB) Delete(value interface{}) *gorm.DB {
	return db.Conn.Delete(value)
}
