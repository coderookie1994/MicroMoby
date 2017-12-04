package controllers

import (
	"net/http"

	"github.com/MicroMoby/data"
)

// GetAllImages :
func GetAllImages(w http.ResponseWriter, r *http.Request) {
	imageList := data.ListAllImages()
	w.Write(imageList)
}

// Hello : test controller
func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
