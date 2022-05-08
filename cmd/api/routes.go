package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/blogs/", app.createBlogHandler)
	router.HandlerFunc(http.MethodGet, "/v1/blogs/:id", app.showBlogHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/blogs/:id", app.updateBlogsHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/blogs/:id", app.deleteBlogHandler)

	return router
}
