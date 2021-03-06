package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteSnapshotResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSnapshotResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteSnapshotResponse struct{}"
	}

	return strings.Join([]string{"DeleteSnapshotResponse", string(data)}, " ")
}
