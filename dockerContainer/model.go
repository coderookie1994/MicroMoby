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
)
