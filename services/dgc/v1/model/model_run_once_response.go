package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RunOnceResponse struct {
	// 作业实例id
	InstanceId     *string `json:"instanceId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunOnceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RunOnceResponse struct{}"
	}

	return strings.Join([]string{"RunOnceResponse", string(data)}, " ")
}
