package main

import (
	"log"
	"net/http"

	"github.com/MicroMoby/routers"
)

func main() {
	router := routers.InitRoutes()
	server := &http.Server{
		Addr:    ":4200",
		Handler: router,
	}
	log.Println("Listening...")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
