package main

import (
	serverConfig "golang-rest-api-starter/internals/config/server"
	"golang-rest-api-starter/router"
	"golang-rest-api-starter/service/middleware"
)

func main() {
	var route = router.NewRouter{
		Middlewares: []router.Middleware{
			middleware.AuthorizedMiddleware,
			middleware.DatabaseMiddleware,
		}}

	// http.HandleFunc("/", handlers.HomeHandler)

	var server = serverConfig.Config{
		PORT:     ":8000",
		Hostname: "http://localhost",
	}
	server.Init(&route)
}
