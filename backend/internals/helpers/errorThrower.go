package helpers

import (
	model "golang-rest-api-starter/models"
	"net/http"
)

func ErrorThrower(err error, message string, statusCode int, w http.ResponseWriter, r *http.Request) {
	response := model.Reponse{
		Message:    message,
		StatusCode: statusCode,
		Data:       nil,
	}

	files := []string{"templates/error.html", "templates/components/base.html"}

	if err != nil {
		ResponseFormatter(
			response, // server response
			files,    // page where to send response
			"error",
			w,
			r,
			statusCode,
		)
		return // Stop execution if there's an error
	}
}
