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

// StartContainerByID starts the container
// matching the container ID given in the
// request body
func (repo *Repository) StartContainerByID(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var rm common.ResponseMessage
	rm.Source = common.ContainerSource

	var startContainerByIDModel StartContainerByIDModel

	if err := json.NewDecoder(r.Body).Decode(&startContainerByIDModel); err != nil {
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
	if err := client.ContainerStart(ctx, startContainerByIDModel.ID, containerStartOptions); err != nil {
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

// StopContainerByID stops the container
// matching the container ID given in the
// request body
func (repo *Repository) StopContainerByID(w http.ResponseWriter, r *http.Request) {
	var rm common.ResponseMessage
	rm.Source = common.ContainerSource

	ctx := context.Background()

	var stopContainerByIDModel StopContainerByIDModel

	if err := json.NewDecoder(r.Body).Decode(&stopContainerByIDModel); err != nil {
		rm.IsSuccess = false
		rm.Message = common.ClientErrorMessage
		rmJSON, _ := json.Marshal(rm)
		log.Println(err)
		w.Write(rmJSON)
		return
	}

	client := gc.Get(r, "dockerClient").(*client.Client)
	if err := client.ContainerStop(ctx, stopContainerByIDModel.ID, nil); err != nil {
		rm.IsSuccess = false
		rm.Message = common.ClientErrorMessage
		rmJSON, _ := json.Marshal(rm)
		log.Println(err)
		w.Write(rmJSON)
		return
	}

	rm.IsSuccess = true
	rm.Message = common.StopContainerSuccessMessage
	rmJSON, _ := json.Marshal(rm)
	w.Write(rmJSON)
}

// ListContainers lists all the containers
// that the docker client connects to
func (repo *Repository) ListContainers(w http.ResponseWriter, r *http.Request) {
	var rm common.ResponseMessage
	rm.Source = common.ContainerSource

	ctx := context.Background()

	options := types.ContainerListOptions{
		All: true,
	}
	client := gc.Get(r, "dockerClient").(*client.Client)
	list, err := client.ContainerList(ctx, options)
	if err != nil {
		rm.IsSuccess = false
		rm.Message = common.ClientErrorMessage
		rmJSON, _ := json.Marshal(rm)
		log.Panicln(err)
		w.Write(rmJSON)
		return
	}

	jsonResponse, err := json.Marshal(buildListResponse(list))
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

func buildListResponse(list []types.Container) []ListContainerResponseModel {
	var listContainerResModel []ListContainerResponseModel
	for _, eachContainer := range list {
		var tempListContainerResModel ListContainerResponseModel
		var ports []port
		tempListContainerResModel.ID = eachContainer.ID
		tempListContainerResModel.Names = eachContainer.Names
		tempListContainerResModel.Image = eachContainer.Image
		tempListContainerResModel.ImageID = eachContainer.ImageID
		tempListContainerResModel.Command = eachContainer.Command
		tempListContainerResModel.State = eachContainer.State
		tempListContainerResModel.Status = eachContainer.Status
		for _, eachPort := range eachContainer.Ports {
			var tempPort port
			tempPort.IP = eachPort.IP
			tempPort.PrivatePort = eachPort.PrivatePort
			tempPort.PublicPort = eachPort.PublicPort
			tempPort.Type = eachPort.Type
			ports = append(ports, tempPort)
		}
		tempListContainerResModel.Ports = ports
		listContainerResModel = append(listContainerResModel, tempListContainerResModel)
	}
	return listContainerResModel
}
