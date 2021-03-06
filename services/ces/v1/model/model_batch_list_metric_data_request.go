package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchListMetricDataRequest struct {
	ContentType string                          `json:"Content-Type"`
	Body        *BatchListMetricDataRequestBody `json:"body,omitempty"`
}

func (o BatchListMetricDataRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchListMetricDataRequest struct{}"
	}

	return strings.Join([]string{"BatchListMetricDataRequest", string(data)}, " ")
}
