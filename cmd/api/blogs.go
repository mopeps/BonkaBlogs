package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mopeps/greenlight/internal/data"
	"github.com/mopeps/greenlight/internal/validator"
)

func (app *application) createBlogHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string   `json:"title"`
		Tags  []string `json:"tags"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	blog := &data.Blog{
		CreatedAt: time.Now(),
		Title:     input.Title,
		Tags:      input.Tags,
	}
	v := validator.New()

	if data.ValidateBlog(v, blog); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Blogs.Insert(blog)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/blogs/%d", blog.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"blog": blog}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showBlogHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	blog := data.Blog{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "John Doe's adventure",
		Tags:      []string{"John", "JOJO", "Doe"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"blog": blog}, nil)
	if err != nil {
		app.logger.Println(err)
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) indexBlogsHandler(w http.ResponseWriter, r *http.Request) {
}

func (app *application) deleteBlogHandler(w http.ResponseWriter, r *http.Request) {
}
