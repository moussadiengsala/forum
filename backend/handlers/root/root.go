package rootHandlers

import (
	"golang-rest-api-starter/internals/helpers"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"templates/components/base.html",
		"templates/home.html",
	}

	helpers.ResponseFormatter(nil, files, "home", w, r, http.StatusOK)
	return
}
