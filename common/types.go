package common

import "net/http"

type (

	// Route :
	Route struct {
		Name        string
		Method      string
		Pattern     string
		HandlerFunc http.HandlerFunc
	}

	// Routes :
	Routes []Route

	// ResponseMessage :
	ResponseMessage struct {
		Source    string `json:"source"`
		IsSuccess bool   `json:"isSuccess"`
		Message   string `json:"message"`
	}
)

const (
	// ContainerSource :
	ContainerSource = "Container"
	// ImageSource :
	ImageSource = "Image"
	// ClientErrorMessage :
	ClientErrorMessage = "something went wrong."
	// StartContainerSuccessMessage :
	StartContainerSuccessMessage = "successfully started container"
	// ListContainerSuccessMessage :
	ListContainerSuccessMessage = "successfully obtained container list"
)
