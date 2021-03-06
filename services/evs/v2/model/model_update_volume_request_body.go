package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type UpdateVolumeRequestBody struct {
	Volume *UpdateVolumeOption `json:"volume"`
}

func (o UpdateVolumeRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateVolumeRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateVolumeRequestBody", string(data)}, " ")
}
