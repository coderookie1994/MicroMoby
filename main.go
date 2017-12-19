package main

import (
	"log"
	"net/http"

	"github.com/MicroMoby/routers"
	"github.com/gorilla/handlers"
)

func main() {
	router := routers.InitRoutes()
	allowedMethods := []string{"GET", "POST", "PUT", "HEAD"}
	// Fix this with a specific url later
	allowedOrigins := []string{"*"}
	server := &http.Server{
		Addr:    ":8808",
		Handler: handlers.CORS(handlers.AllowedMethods(allowedMethods), handlers.AllowedOrigins(allowedOrigins))(router),
	}
	log.Println("Listening...")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
