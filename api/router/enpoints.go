package router

import (
	authHandlers "golang-rest-api-starter/handlers/auth"
	postHandlers "golang-rest-api-starter/handlers/posts"
	rootHandlers "golang-rest-api-starter/handlers/root"
	usersHandlers "golang-rest-api-starter/handlers/users"
	"golang-rest-api-starter/internals/helpers"
	model "golang-rest-api-starter/models"
	"net/http"
	"strings"
)

func Handlers(w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	var routeHandlers = map[string]http.HandlerFunc{
		"/":              rootHandlers.HomeHandler,
		"/posts":         postHandlers.Posts,
		"/posts/":        postHandlers.Post,
		"/users/":        usersHandlers.Users,
		"/auth/register": authHandlers.Register,
		"/auth/logout":   authHandlers.LogOut,
		"/auth/login":    authHandlers.LogIn,
	}

	// Handle routes with specific prefixes
	for prefix, handlerFunc := range routeHandlers {
		if r.URL.Path == prefix || strings.HasPrefix(r.URL.Path, prefix) && helpers.Contains([]string{"/posts/", "/users/"}, prefix) {
			return handlerFunc
		}
	}

	return func(w http.ResponseWriter, _ *http.Request) {
		var response = model.Reponse{
			Message:    "Oops! This page that you looking for does not exist. It might be move or delete",
			StatusCode: http.StatusNotFound,
			Data:       nil,
		}
		helpers.ResponseFormatter(response, w, r, http.StatusNotFound)
	}
}
