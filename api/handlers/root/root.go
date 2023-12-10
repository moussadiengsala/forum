package rootHandlers

import (
	"golang-rest-api-starter/internals/helpers"
	model "golang-rest-api-starter/models"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	var data = model.Reponse{
		Message:    "OK",
		Data:       "Hello, world api in galang.com",
		StatusCode: 200,
	}

	helpers.ResponseFormatter(data, w, r, http.StatusOK)
	return
}
