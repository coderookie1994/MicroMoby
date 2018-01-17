package main

import (
	"log"
	"net/http"

	"github.com/MicroMoby/dockerImage"
	"github.com/gorilla/handlers"
)

func main() {
	router := dockerImage.NewRouter()
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD"})
	// Fix this with a specific url later
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	server := &http.Server{
		Addr:    ":8808",
		Handler: handlers.CORS(allowedMethods, allowedOrigins)(router),
	}
	log.Println("Listening...")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
