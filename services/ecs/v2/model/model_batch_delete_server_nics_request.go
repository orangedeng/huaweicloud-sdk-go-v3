package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchDeleteServerNicsRequest struct {
	ServerId string                            `json:"server_id"`
	Body     *BatchDeleteServerNicsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteServerNicsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchDeleteServerNicsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteServerNicsRequest", string(data)}, " ")
}
