package helpers

import (
	"log"
	"net/http"
	"text/template"
)

/*
This function has as purpose formatted the
response from the server & send it the client
*/
func ResponseFormatter(data any, files []string, page string, w http.ResponseWriter, r *http.Request, statusCode int) {
	w.WriteHeader(statusCode)
	log.Println("Request URL:", r.URL.Path, "Status Code:", statusCode)

	tmpl := template.Must(template.New("").ParseFiles(files...))

	tmpl.ExecuteTemplate(w, page, data)
	return
}
