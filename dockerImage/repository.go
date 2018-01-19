package dockerImage

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

// ListAllImages is a helper method that
// uses the client object passed in the
// request context and gets the list of
// all available docker images that the
// client object is connected to
func (repo *Repository) ListAllImages(w http.ResponseWriter, r *http.Request) {
	var rm common.ResponseMessage
	rm.Source = common.ImageSource

	var imageList []ImageResponseModel
	client := gc.Get(r, "dockerClient").(*client.Client)

	images, err := client.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		rm.IsSuccess = false
		rm.Message = common.ClientErrorMessage
		rmJSON, _ := json.Marshal(rm)
		log.Println(err)
		w.Write(rmJSON)
		return
	}

	for _, image := range images {
		var tempImagelist ImageResponseModel
		tempImagelist.ID = image.ID
		tempImagelist.Containers = image.Containers
		tempImagelist.Size = image.Size
		tempImagelist.RepoTags = image.RepoTags
		imageList = append(imageList, tempImagelist)
	}

	imageListJSON, err := json.Marshal(imageList)
	if err != nil {
		rm.IsSuccess = false
		rm.Message = common.ClientErrorMessage
		rmJSON, _ := json.Marshal(rm)
		log.Println(err)
		w.Write(rmJSON)
		return
	}

	rm.IsSuccess = true
	w.Write(imageListJSON)
}
