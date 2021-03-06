package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type ResizeVolumeRequestBody struct {
	BssParam *BssParamForResizeVolume `json:"bssParam,omitempty"`
	OsExtend *OsExtend                `json:"os-extend"`
}

func (o ResizeVolumeRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ResizeVolumeRequestBody struct{}"
	}

	return strings.Join([]string{"ResizeVolumeRequestBody", string(data)}, " ")
}
