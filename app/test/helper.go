package test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBMock creates an SQL mock object, and associated gorm.DB object.
func DBMock(t *testing.T) (sqlmock.Sqlmock, *gorm.DB) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf(`sqlmock.New()=_, _, %#v; want nil`, err)
	}

	gormDB, err := gorm.Open(mysql.Dialector{Config: &mysql.Config{DriverName: "mysql", Conn: db, SkipInitializeWithVersion: true}}, &gorm.Config{})

	if err != nil {
		t.Fatalf(`gorm.Open("mysql", %#v)=_, %#v; want nil`, db, err)
	}

	gormDB.Debug()

	return mock, gormDB
}
