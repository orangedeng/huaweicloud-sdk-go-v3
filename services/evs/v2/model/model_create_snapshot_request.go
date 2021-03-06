package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateSnapshotRequest struct {
	Body *CreateSnapshotRequestBody `json:"body,omitempty"`
}

func (o CreateSnapshotRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateSnapshotRequest struct{}"
	}

	return strings.Join([]string{"CreateSnapshotRequest", string(data)}, " ")
}
