package data

import (
	"time"

	"github.com/mopeps/greenlight/internal/validator"
)

type Blog struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Tags      []string  `json:"tags,omitempty"`
	Version   int32     `json:"version"`
}

func ValidateBlog(v *validator.Validator, blog *Blog) {
	v.Check(blog.Title != "", "title", "Must provide one")
	v.Check(len(blog.Title) <= 500, "title", "length shouldn't be superior to 500 bytes")

	v.Check(blog.Tags != nil, "tags", "must provide one") // i 'll have to think this one
	v.Check(len(blog.Tags) <= 20, "tags", "length shouldn't be superior to 20")
}
