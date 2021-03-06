package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListInstanceMetricRequest struct {
	BlockchainId string                         `json:"blockchain_id"`
	Body         *ListInstanceMetricRequestBody `json:"body,omitempty"`
}

func (o ListInstanceMetricRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListInstanceMetricRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceMetricRequest", string(data)}, " ")
}
