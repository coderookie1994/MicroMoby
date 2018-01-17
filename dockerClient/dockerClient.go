package dockerClient

import (
	"log"
	"net/http"

	"github.com/MicroMoby/adapter"
	"github.com/docker/docker/client"
	"github.com/gorilla/context"
)

// DockerClient :
func DockerClient() adapter.HandlerAdapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cli, err := client.NewEnvClient()
			if err != nil {
				log.Fatalln(err)
			}
			defer cli.Close()

			context.Set(r, "dockerClient", cli)
			h.ServeHTTP(w, r)
		})
	}
}
