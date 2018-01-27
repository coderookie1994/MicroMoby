package dockerContainer

import (
	"net/http"
)

// Controller :
type Controller struct {
	Repository Repository
}

// StartContainer :
func (c *Controller) StartContainer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	c.Repository.StartContainerByID(w, r)
}

// StopContainer :
func (c *Controller) StopContainer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	c.Repository.StopContainerByID(w, r)
}

// ListAllContainers :
func (c *Controller) ListAllContainers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	c.Repository.ListContainers(w, r)
}
