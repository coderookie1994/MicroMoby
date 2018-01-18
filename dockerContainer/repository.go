package dockerContainer

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	gc "github.com/gorilla/context"
)

// Repository :
type Repository struct {
}

// StartContainerByID : start the container from the
func (repo *Repository) StartContainerByID(w http.ResponseWriter, r *http.Request) {
	var startContainerByIDModel StartContainerByIDModel
	err := json.NewDecoder(r.Body).Decode(&startContainerByIDModel)
	if err != nil {
		log.Println(err)
		w.Write([]byte("something went wrong"))
		return
	}

	client := gc.Get(r, "dockerClient").(*client.Client)
	var containerStartOptions types.ContainerStartOptions
	containerStartOptions.CheckpointID = startContainerByIDModel.CheckpointID
	containerStartOptions.CheckpointDir = startContainerByIDModel.CheckpointDir
	err = client.ContainerStart(context.Background(), startContainerByIDModel.ID, containerStartOptions)
	if err != nil {
		log.Println(err)
		w.Write([]byte("something went wrong"))
		return
	}
	w.Write([]byte("container successfully started"))
}
