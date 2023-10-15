package router

import (
	"golang-rest-api-starter/handlers"
	"net/http"
	"strings"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc
type NewRouter struct {
	Middlewares []Middleware
}

func (router *NewRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	handler := func(w http.ResponseWriter, r *http.Request) {
		switch {
		// Home Page
		case r.URL.Path == "/":
			handlers.HomeHandler(w, r)
		// Static files
		case strings.HasPrefix(r.URL.Path, "/static/"):
			fs := http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))
			fs.ServeHTTP(w, r)

		// When the URL doesn't exist
		default:
		}
	}

	for _, mw := range router.Middlewares {
		handler = mw(handler)
	}
	handler(w, r)
}
