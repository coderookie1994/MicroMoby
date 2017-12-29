package data

import (
	"golang.org/x/net/context"

	"github.com/MicroMoby/models"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// StartContainerID :
func StartContainerID(startContainerByIDModel models.StartContainerByIDModel) bool {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	defer cli.Close()
	var containerStartOptions types.ContainerStartOptions
	containerStartOptions.CheckpointID = startContainerByIDModel.CheckpointID
	containerStartOptions.CheckpointDir = startContainerByIDModel.CheckpointDir
	err = cli.ContainerStart(context.Background(), startContainerByIDModel.ID, containerStartOptions)
	if err != nil {
		// handle this err: log it
		return false
	}
	return true
}
