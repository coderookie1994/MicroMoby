package models

// ImageResponseModel :
type (
	ImageResponseModel struct {
		ID         string   `json:"id"`
		Containers int64    `json:"containers"`
		Size       int64    `json:"size"`
		RepoTags   []string `json:"repoTags"`
	}
)
