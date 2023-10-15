package handlers

import (
	components "golang-rest-api-starter/Components"
	"net/http"
	"text/template"
)

var data = map[string]interface{}{
	"name": "Moussa",
	"sex":  "Male",
	"age":  75,
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var funcM = template.FuncMap{
		"test": components.Test,
	}

	tmpl := template.Must(template.New("").Funcs(funcM).ParseFiles("templates/base.html", "templates/home.html"))

	tmpl.ExecuteTemplate(w, "home.html", data)
	return
}
