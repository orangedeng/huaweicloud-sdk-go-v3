package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type UpdateSnapshotRequestBody struct {
	Snapshot *UpdateSnapshotOption `json:"snapshot"`
}

func (o UpdateSnapshotRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateSnapshotRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSnapshotRequestBody", string(data)}, " ")
}
