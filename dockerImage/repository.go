package dockerImage

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

// ListAllImages is a helper method that
// uses the client object passed in the
// request context and gets the list of
// all available docker images that the
// client object is connected to
func (repo *Repository) ListAllImages(w http.ResponseWriter, r *http.Request) {
	var imageList []ImageResponseModel
	client := gc.Get(r, "dockerClient").(*client.Client)

	images, err := client.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		log.Println(err)
		w.Write([]byte("something went wrong"))
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
		log.Println(err)
		w.Write([]byte("something went wrong"))
		return
	}
	w.Write(imageListJSON)
}
