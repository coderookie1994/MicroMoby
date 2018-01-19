package dockerContainer

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/MicroMoby/common"
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

	var rm common.ResponseMessage
	rm.Source = common.ContainerSource

	err := json.NewDecoder(r.Body).Decode(&startContainerByIDModel)
	if err != nil {
		rm.IsSuccess = false
		rm.Message = common.ClientErrorMessage
		rmJSON, _ := json.Marshal(rm)
		log.Println(err)
		w.Write(rmJSON)
		return
	}

	client := gc.Get(r, "dockerClient").(*client.Client)
	var containerStartOptions types.ContainerStartOptions
	containerStartOptions.CheckpointID = startContainerByIDModel.CheckpointID
	containerStartOptions.CheckpointDir = startContainerByIDModel.CheckpointDir
	// nil check for containerStartOptions props and send
	err = client.ContainerStart(context.Background(), startContainerByIDModel.ID, containerStartOptions)
	if err != nil {
		rm.IsSuccess = false
		rm.Message = common.ClientErrorMessage
		rmJSON, _ := json.Marshal(rm)
		log.Println(err)
		w.Write(rmJSON)
		return
	}
	rm.IsSuccess = true
	rm.Message = common.StartContainerSuccessMessage
	rmJSON, _ := json.Marshal(rm)
	w.Write(rmJSON)
}

// ListContainers :
func (repo *Repository) ListContainers(w http.ResponseWriter, r *http.Request) {
	var rm common.ResponseMessage
	rm.Source = common.ContainerSource

	options := types.ContainerListOptions{
		All: true,
	}
	client := gc.Get(r, "dockerClient").(*client.Client)
	list, err := client.ContainerList(context.Background(), options)
	if err != nil {
		rm.IsSuccess = false
		rm.Message = common.ClientErrorMessage
		rmJSON, _ := json.Marshal(rm)
		log.Panicln(err)
		w.Write(rmJSON)
		return
	}

	jsonResponse, err := json.Marshal(list)
	if err != nil {
		rm.IsSuccess = false
		rm.Message = common.ClientErrorMessage
		rmJSON, _ := json.Marshal(rm)
		log.Panicln(err)
		w.Write(rmJSON)
		return
	}

	rm.IsSuccess = true
	w.Write(jsonResponse)
}
