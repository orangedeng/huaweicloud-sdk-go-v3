package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowCheckpointResponse struct {
	Checkpoint     *CheckpointCreate `json:"checkpoint,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowCheckpointResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowCheckpointResponse struct{}"
	}

	return strings.Join([]string{"ShowCheckpointResponse", string(data)}, " ")
}
