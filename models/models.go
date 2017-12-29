package models

type (
	// ImageResponseModel :
	ImageResponseModel struct {
		ID         string   `json:"id"`
		Containers int64    `json:"containers"`
		Size       int64    `json:"size"`
		RepoTags   []string `json:"repoTags"`
	}

	// StartContainerByIDModel :
	StartContainerByIDModel struct {
		ID            string `json:"id"`
		CheckpointID  string `json:"checkpointId"`
		CheckpointDir string `json:"checkpointDir"`
	}
)
