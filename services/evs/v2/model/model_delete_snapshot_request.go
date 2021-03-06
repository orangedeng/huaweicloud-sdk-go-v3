package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteSnapshotRequest struct {
	SnapshotId string `json:"snapshot_id"`
}

func (o DeleteSnapshotRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteSnapshotRequest struct{}"
	}

	return strings.Join([]string{"DeleteSnapshotRequest", string(data)}, " ")
}
