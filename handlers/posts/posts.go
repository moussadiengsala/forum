package postHandlers

import (
	"golang-rest-api-starter/internals/helpers"
	"net/http"
)

func Posts(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/components/base.html",
		"templates/components/header.html",
		"templates/components/aside.html",
		"templates/posts/posts.html",
	}

	helpers.ResponseFormatter(nil, files, "posts", w, r, http.StatusOK)
	return
}
