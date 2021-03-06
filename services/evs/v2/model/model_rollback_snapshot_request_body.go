package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type RollbackSnapshotRequestBody struct {
	Rollback *RollbackSnapshotOption `json:"rollback"`
}

func (o RollbackSnapshotRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RollbackSnapshotRequestBody struct{}"
	}

	return strings.Join([]string{"RollbackSnapshotRequestBody", string(data)}, " ")
}
