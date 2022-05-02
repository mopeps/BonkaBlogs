package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNortFound = errors.New("record not found")
)

type Models struct {
	Blogs BlogModel
	Blog  interface {
		Insert(blog *Blog) error
		Get(id int64) (*Blog, error)
		Update(blog *Blog) error
		Delete(id int64) error
	}
}

func NewModels(db *sql.DB) Models {
	return Models{
		Blogs: BlogModel{DB: db},
	}
}

func NewMockModels(db *sql.DB) Models {
	return Models{
		Blogs: MockBlogModel{},
	}
}
