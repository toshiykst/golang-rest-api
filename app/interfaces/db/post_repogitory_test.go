package db

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
	"gorm.io/gorm"

	"github.com/toshiykst/golang-rest-api/app/domain"
	"github.com/toshiykst/golang-rest-api/app/test"
)

// MockDB is a structure of the MockDB.
type MockDB struct {
	Conn *gorm.DB
}

// Find returns the record by query.
func (db *MockDB) Find(out interface{}, where ...interface{}) *gorm.DB {
	return db.Conn.Find(out, where...)
}

// Create creates the new record.
func (db *MockDB) Create(value interface{}) *gorm.DB {
	return db.Conn.Create(value)
}

// Update updates the record.
func (db *MockDB) Update(value interface{}) *gorm.DB {
	return db.Conn.Save(value)
}

// Delete deletes the record.
func (db *MockDB) Delete(value interface{}) *gorm.DB {
	return db.Conn.Delete(value)
}

func TestPostRepository_FindPost(t *testing.T) {
	mock, conn := test.DBMock(t)
	sqlDB, err := conn.DB()
	defer sqlDB.Close()

	if err != nil {
		t.Fatalf("want no err, but has error %#v", err)
	}

	want := domain.Post{ID: 1, Title: "title", Content: "content"}

	id := 1

	mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `posts` WHERE `posts`.`id` = ?")).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "content"}).
			AddRow(want.ID, want.Title, want.Content))

	r := &PostRepository{DB: &MockDB{Conn: conn}}

	got, err := r.FindPost(id)

	if err != nil {
		t.Fatalf("want no err, but has error %#v", err)
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("TestPostRepository_FindPost differs: (-got +want)\n%s", diff)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPostRepository_CreatePost(t *testing.T) {
	mock, conn := test.DBMock(t)
	sqlDB, err := conn.DB()
	defer sqlDB.Close()

	if err != nil {
		t.Fatalf("want no err, but has error %#v", err)
	}

	post := domain.Post{Title: "title", Content: "content"}

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `posts` (`title`,`content`,`created_at`,`updated_at`) VALUES (?,?,?,?)")).
		WithArgs(post.Title, post.Content, test.AnyTime{}, test.AnyTime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))

	r := &PostRepository{DB: &MockDB{Conn: conn}}

	if err = r.CreatePost(&post); err != nil {
		t.Fatalf("want no err, but has error %#v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPostRepository_UpdatePost(t *testing.T) {
	mock, conn := test.DBMock(t)
	sqlDB, err := conn.DB()
	defer sqlDB.Close()

	if err != nil {
		t.Fatalf("want no err, but has error %#v", err)
	}

	post := domain.Post{ID: 1, Title: "updated title", Content: "updated content"}

	mock.ExpectExec(regexp.QuoteMeta("UPDATE `posts` SET `title`=?,`content`=?,`created_at`=?,`updated_at`=? WHERE `id` = ?")).
		WithArgs(post.Title, post.Content, nil, test.AnyTime{}, post.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	r := &PostRepository{DB: &MockDB{Conn: conn}}

	if err = r.UpdatePost(&post); err != nil {
		t.Fatalf("want no err, but has error %#v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPostRepository_DeletePost(t *testing.T) {
	mock, conn := test.DBMock(t)
	sqlDB, err := conn.DB()
	defer sqlDB.Close()

	if err != nil {
		t.Fatalf("want no err, but has error %#v", err)
	}

	post := domain.Post{ID: 1, Title: "title", Content: "content"}

	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `posts` WHERE `posts`.`id` = ?")).
		WithArgs(post.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	r := &PostRepository{DB: &MockDB{Conn: conn}}

	if err = r.DeletePost(&post); err != nil {
		t.Fatalf("want no err, but has error %#v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPostRepository_FindPosts(t *testing.T) {
	mock, conn := test.DBMock(t)
	sqlDB, err := conn.DB()
	defer sqlDB.Close()

	if err != nil {
		t.Fatalf("want no err, but has error %#v", err)
	}

	want := domain.Posts{
		{ID: 1, Title: "title1", Content: "content1"},
		{ID: 2, Title: "title2", Content: "content2"},
		{ID: 3, Title: "title3", Content: "content3"},
	}

	rows := sqlmock.NewRows([]string{"id", "title", "content"})
	for _, p := range want {
		rows.AddRow(p.ID, p.Title, p.Content)
	}

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `posts`")).WillReturnRows(rows)

	r := &PostRepository{DB: &MockDB{Conn: conn}}

	got, err := r.FindPosts()

	if err != nil {
		t.Fatalf("want no err, but has error %#v", err)
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("TestPostRepository_FindPosts differs: (-got +want)\n%s", diff)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
