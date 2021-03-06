package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateCheckpointResponse struct {
	Checkpoint     *CheckpointCreate `json:"checkpoint,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o CreateCheckpointResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateCheckpointResponse struct{}"
	}

	return strings.Join([]string{"CreateCheckpointResponse", string(data)}, " ")
}
