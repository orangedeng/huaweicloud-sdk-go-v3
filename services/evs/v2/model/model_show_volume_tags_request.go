package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowVolumeTagsRequest struct {
	VolumeId string `json:"volume_id"`
}

func (o ShowVolumeTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowVolumeTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowVolumeTagsRequest", string(data)}, " ")
}
