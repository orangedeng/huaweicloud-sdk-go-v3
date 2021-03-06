package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListTranscodeDetailRequest struct {
	TaskId *[]string `json:"task_id,omitempty"`
}

func (o ListTranscodeDetailRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListTranscodeDetailRequest struct{}"
	}

	return strings.Join([]string{"ListTranscodeDetailRequest", string(data)}, " ")
}
