package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateSnapshotResponse struct {
	Snapshot       *SnapshotDetails `json:"snapshot,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o CreateSnapshotResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateSnapshotResponse struct{}"
	}

	return strings.Join([]string{"CreateSnapshotResponse", string(data)}, " ")
}
