package main

import (
	"fmt"
	"net/http"

	core "learn.zone01dakar.sn/forum-rest-api/internals/core"
	"learn.zone01dakar.sn/forum-rest-api/routes"
	"learn.zone01dakar.sn/forum-rest-api/routes/feedflow"
	service "learn.zone01dakar.sn/forum-rest-api/service/middlewares"
)

func main() {
	// Creating a new instance of our mini-framework
	app := core.NewApp()
	// Middleware usage
	// app.Use(service.CORSMiddleware)
	app.Use(service.LoggerMiddleware)
	app.Use(service.AuthMiddleware)
	app.Use(service.DBMiddleware)

	// Defining the differents endpoints
	var endpoints = []routes.Router{&routes.Auth{}, &feedflow.Comment{}, &feedflow.Post{}, &feedflow.Reaction{}}
	routes.Handle(endpoints, app)

	fmt.Println("server runnin at http://localhost:8000")
	// Serving our app
	http.ListenAndServe(":8000", app)
}
