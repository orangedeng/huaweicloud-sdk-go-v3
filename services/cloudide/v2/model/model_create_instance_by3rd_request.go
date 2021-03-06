package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateInstanceBy3rdRequest struct {
	InstanceLabel *string            `json:"instance_label,omitempty"`
	Body          *InstanceEdgeParam `json:"body,omitempty"`
}

func (o CreateInstanceBy3rdRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateInstanceBy3rdRequest struct{}"
	}

	return strings.Join([]string{"CreateInstanceBy3rdRequest", string(data)}, " ")
}
