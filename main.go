package main

import (
	"log"
	"net/http"

	"github.com/MicroMoby/dockerContainer"
	"github.com/MicroMoby/dockerImage"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router = dockerImage.NewRouter(router)
	router = dockerContainer.NewRouter(router)

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
