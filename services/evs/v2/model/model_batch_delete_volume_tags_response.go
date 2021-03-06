package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchDeleteVolumeTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteVolumeTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchDeleteVolumeTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteVolumeTagsResponse", string(data)}, " ")
}
