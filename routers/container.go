package routers

import (
	"github.com/MicroMoby/controllers"
	"github.com/gorilla/mux"
)

// SetContainerRoutes :
func SetContainerRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/container/start", controllers.StartContainerByID).Methods("GET")
	return router
}
