package dockerContainer

import "time"

type (
	// StartContainerByIDModel represents the
	// model the json request will be decoded
	// to.
	StartContainerByIDModel struct {
		ID            string `json:"id"`
		CheckpointID  string `json:"checkpointId"`
		CheckpointDir string `json:"checkpointDir"`
	}

	// StopContainerByIDModel represents the
	// model the json request will be decoded
	// to.
	StopContainerByIDModel struct {
		ID       string         `json:"id"`
		Duration *time.Duration `json:"duration"`
	}

	// ListContainerResponseModel represent the
	// model that will be sent to the client
	// on request of a List of containers
	// for the connected docker daemon.
	ListContainerResponseModel struct {
		ID      string   `json:"id"`
		Names   []string `json:"names"`
		Image   string   `json:"image"`
		ImageID string   `json:"imageId"`
		Command string   `json:"command"`
		Ports   []port   `json:"ports"`
		State   string   `json:"state"`
		Status  string   `json:"status"`
	}

	port struct {
		IP          string `json:"ip"`
		PrivatePort uint16 `json:"privatePort"`
		PublicPort  uint16 `json:"publicPort"`
		Type        string `json:"type"`
	}
)
