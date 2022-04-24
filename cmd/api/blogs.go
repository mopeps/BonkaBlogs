package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mopeps/greenlight/internal/data"
)

func (app *application) createBlogHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new blog")
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
