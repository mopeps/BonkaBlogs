package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
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
