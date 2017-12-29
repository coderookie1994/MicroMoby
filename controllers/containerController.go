package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/MicroMoby/data"

	"github.com/MicroMoby/models"
)

// StartContainerByID : start the container from the
func StartContainerByID(w http.ResponseWriter, r *http.Request) {
	var startContainerByIDModel models.StartContainerByIDModel
	err := json.NewDecoder(r.Body).Decode(&startContainerByIDModel)
	if err != nil {
		panic(err)
	}
	if data.StartContainerID(startContainerByIDModel) {
		w.Write([]byte("Container start success message"))
	} else {
		w.Write([]byte("Container start failed message"))
	}
}
