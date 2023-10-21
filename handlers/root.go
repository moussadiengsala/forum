package handlers

import (
	"net/http"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"templates/base.html",
		"templates/home.html",
		"templates/components/header.html",
		"templates/components/aside.html",
	}

	tmpl := template.Must(template.New("").ParseFiles(files...))

	tmpl.ExecuteTemplate(w, "base", nil)
	return
}
