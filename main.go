package main

import (
	"log"
	"net/http"

	"github.com/MicroMoby/routers"
	"github.com/gorilla/handlers"
)

func main() {
	router := routers.InitRoutes()
	// server := &http.Server{
	// 	Addr:    ":4200",
	// 	Handler: router,
	// }
	log.Println("Listening...")
	err := http.ListenAndServe(":8808", handlers.CORS(handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD"}), handlers.AllowedOrigins([]string{"*"}))(router))
	if err != nil {
		panic(err)
	}
}
