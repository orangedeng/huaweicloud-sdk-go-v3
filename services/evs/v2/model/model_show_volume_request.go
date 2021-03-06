package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowVolumeRequest struct {
	VolumeId string `json:"volume_id"`
}

func (o ShowVolumeRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowVolumeRequest struct{}"
	}

	return strings.Join([]string{"ShowVolumeRequest", string(data)}, " ")
}
