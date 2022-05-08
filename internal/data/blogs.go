package data

import (
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
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

type BlogModel struct {
	DB *sql.DB
}

func (m BlogModel) Insert(blog *Blog) error {
	query := `
		INSERT INTO blogs (title, tags)
		VALUES($1, $2)
		RETURNING id, created_at, version`
	args := []interface{}{blog.Title, pq.Array(blog.Tags)}
	return m.DB.QueryRow(query, args...).Scan(&blog.ID, &blog.CreatedAt, &blog.Version)
}

func (m BlogModel) Get(id int64) (*Blog, error) {

	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `
		SELECT id, created_at, title, tags, version
		FROM blogs
		WHERE id = $1`

	var blog Blog

	err := m.DB.QueryRow(query, id).Scan(
		&blog.ID,
		&blog.CreatedAt,
		&blog.Title,
		pq.Array(&blog.Tags),
		&blog.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &blog, nil
}
func (m BlogModel) Update(blog *Blog) error {
	query := `UPDATE blogs
		SET title = $1, tags = $2, version = version + 1
		WHERE id = $3
		RETURNING version`

	args := []interface{}{
		blog.Title,
		pq.Array(blog.Tags),
	}
	return m.DB.QueryRow(query, args...).Scan(&blog.Version)
}

func (m BlogModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := ` DELETE FROM blogs
		WHERE id = $1`

	result, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}

type MockBlogModel struct{}

func (m MockBlogModel) Insert(blog *Blog) error {
	return nil
}

func (m MockBlogModel) Get(id int64) (*Blog, error) {
	return nil, nil
}
func (m MockBlogModel) Update(blog *Blog) error {
	return nil
}

func (m MockBlogModel) Delete(id int64) error {
	return nil
}
