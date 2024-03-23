package main

import (
	server "golang-rest-api-starter/internals/config/server"
	"golang-rest-api-starter/router"
	"golang-rest-api-starter/service/middleware"
	"log"
	"os"

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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// instantiate a new server
	server := server.Config{
		PORT:     port,
		Hostname: "http://localhost",
	}
	server.Init(router)
}
