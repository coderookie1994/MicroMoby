package dockerContainer

import (
	"net/http"

	"github.com/MicroMoby/adapter"
	"github.com/MicroMoby/common"
	"github.com/MicroMoby/dockerClient"
	"github.com/gorilla/mux"
)

var controller = &Controller{Repository: Repository{}}

var routes = common.Routes{
	common.Route{
		Name:        "ListContainers",
		Method:      "GET",
		Pattern:     "/listcontainer",
		HandlerFunc: controller.ListAllContainers,
	},
	common.Route{
		Name:        "StartContainer",
		Method:      "POST",
		Pattern:     "/startcontainer",
		HandlerFunc: controller.StartContainer,
	},
	common.Route{
		Name:        "StopContainer",
		Method:      "POST",
		Pattern:     "/stopcontainer",
		HandlerFunc: controller.StopContainer,
	},
}

// NewRouter :
func NewRouter(router *mux.Router) *mux.Router {
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = adapter.AdaptHandler(handler, dockerClient.DockerClient())

		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}
	return router
}
