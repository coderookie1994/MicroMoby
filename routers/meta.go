package routers

import (
	"github.com/MicroMoby/controllers"
	"github.com/gorilla/mux"
)

// SetMetaRoutes :
func SetMetaRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/meta/images", controllers.GetAllImages).Methods("GET")
	router.HandleFunc("/meta/hello", controllers.Hello).Methods("GET")
	return router
}
