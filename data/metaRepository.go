package data

import (
	"encoding/json"

	"github.com/MicroMoby/models"

	"golang.org/x/net/context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// ListAllImages :
func ListAllImages() []byte {
	var imageList []models.ImageResponseModel

	cli, err := client.NewEnvClient()

	if err != nil {
		// figure out how to handle this..
		panic(err)
	}

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		// figure out how to handle this..
		panic(err)
	}

	for _, image := range images {
		var tempImagelist models.ImageResponseModel
		tempImagelist.ID = image.ID
		tempImagelist.Containers = image.Containers
		tempImagelist.Size = image.Size
		tempImagelist.RepoTags = image.RepoTags
		imageList = append(imageList, tempImagelist)
	}
	imageListJSON, err := json.Marshal(imageList)
	if err != nil {
		panic(err)
	}
	return imageListJSON
}
