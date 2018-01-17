package dockerImage

import (
	"net/http"
)

type Controller struct {
	Repository Repository
}

func (c *Controller) GetAllImages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	c.Repository.ListAllImages(w, r)
}
