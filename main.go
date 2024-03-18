package main

import (
	server "golang-rest-api-starter/internals/config/server"
	"golang-rest-api-starter/router"
	"golang-rest-api-starter/service/middleware"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading godotenv: ", err)
	}

	router := &router.NewRouter{
		Middlewares: []router.Middleware{
			middleware.AuthMiddleware,
			middleware.LoggerMiddleware,
			middleware.DBMiddleware,
		},
	}

	// instantiate a new server
	server := server.Config{
		PORT:     ":8080",
		Hostname: "http://localhost",
	}
	server.Init(router)
}
