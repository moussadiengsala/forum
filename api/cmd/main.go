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

	// Defining the differents endpoints
	var endpoints = []routes.Router{&routes.Auth{}, &feedflow.Comment{}, &feedflow.Post{}, &feedflow.Reaction{}}

	routes.Handle(endpoints, app)

	// Middleware usage
	app.Use(service.LoggerMiddleware)
	app.Use(service.AuthMiddleware)
	app.Use(service.DBMiddleware)
	app.Use(service.CORSMiddleware)

	fmt.Println("server runnin at http://localhost:7000")
	// Serving our app
	http.ListenAndServe(":7000", app)
}
