package dockerImage

import (
	"net/http"

	"github.com/MicroMoby/adapter"
	"github.com/MicroMoby/dockerClient"
	"github.com/gorilla/mux"
)

var controller = &Controller{Repository: Repository{}}

// Route :
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes :
type Routes []Route

var routes = Routes{
	Route{
		Name:        "ImageList",
		Method:      "GET",
		Pattern:     "/images",
		HandlerFunc: controller.GetAllImages,
	},
}

// NewRouter :
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = adapter.AdaptHandler(handler, dockerClient.DockerClient())

		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}
	return router
}
