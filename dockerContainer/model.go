package dockerContainer

type (
	// StartContainerByIDModel :
	StartContainerByIDModel struct {
		ID            string `json:"id"`
		CheckpointID  string `json:"checkpointId"`
		CheckpointDir string `json:"checkpointDir"`
	}
)
