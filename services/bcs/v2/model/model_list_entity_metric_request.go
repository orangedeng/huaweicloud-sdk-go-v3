package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListEntityMetricRequest struct {
	BlockchainId string                       `json:"blockchain_id"`
	Body         *ListEntityMetricRequestBody `json:"body,omitempty"`
}

func (o ListEntityMetricRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListEntityMetricRequest struct{}"
	}

	return strings.Join([]string{"ListEntityMetricRequest", string(data)}, " ")
}
