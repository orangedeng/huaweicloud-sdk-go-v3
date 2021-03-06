package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteBackgroundTaskRequest struct {
	InstanceId string `json:"instance_id"`
	TaskId     string `json:"task_id"`
}

func (o DeleteBackgroundTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteBackgroundTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteBackgroundTaskRequest", string(data)}, " ")
}
