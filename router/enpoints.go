package router

import (
	"errors"
	postHandlers "golang-rest-api-starter/handlers/posts"
	rootHandlers "golang-rest-api-starter/handlers/root"
	"golang-rest-api-starter/internals/helpers"
	"net/http"
	"strings"
)

func Handlers(w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	var endpoints = map[string]http.HandlerFunc{
		"/":      rootHandlers.HomeHandler,
		"/posts": postHandlers.Posts,
	}

	if strings.HasPrefix(r.URL.Path, "/static/") {
		fs := http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))
		return fs.ServeHTTP // Set the handler to nil and return true to indicate the path was handled
	}

	var endpoint, ok = endpoints[r.URL.Path]

	if !ok {
		return func(w http.ResponseWriter, _ *http.Request) {
			helpers.ErrorThrower(errors.New("Not Found"), "Oops! This page that you looking for does not exist. It might be move or delete", http.StatusNotFound, w, r)
		}
	}

	return endpoint
}
