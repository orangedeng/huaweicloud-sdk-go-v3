package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteVolumeResponse struct {
	// 正常返回时返回的任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteVolumeResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteVolumeResponse struct{}"
	}

	return strings.Join([]string{"DeleteVolumeResponse", string(data)}, " ")
}
