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
	// nil check for containerStartOptions props and send
	err = client.ContainerStart(context.Background(), startContainerByIDModel.ID, containerStartOptions)
	if err != nil {
		log.Println(err)
		w.Write([]byte("something went wrong"))
		return
	}
	w.Write([]byte("container successfully started"))
}

// ListContainers :
func (repo *Repository) ListContainers(w http.ResponseWriter, r *http.Request) {
	options := types.ContainerListOptions{
		All: true,
	}
	client := gc.Get(r, "dockerClient").(*client.Client)
	list, err := client.ContainerList(context.Background(), options)
	if err != nil {
		log.Panicln([]byte("something went wrong"))
		w.Write([]byte("unable to get all containers"))
		return
	}
	jsonResponse, err := json.Marshal(list)
	if err != nil {
		log.Panicln([]byte("something went wrong"))
		w.Write([]byte("unable to transform response"))
		return
	}
	w.Write(jsonResponse)
}
