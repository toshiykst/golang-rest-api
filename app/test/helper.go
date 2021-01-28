package test

import (
	"database/sql/driver"
	"testing"
	"time"

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

// AnyTime is a structure of AnyTime.
type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}
