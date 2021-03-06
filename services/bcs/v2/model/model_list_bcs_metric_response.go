package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListBcsMetricResponse struct {
	// 指标对象列表。
	Metrics        *[]MetricItemResultApi `json:"metrics,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListBcsMetricResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListBcsMetricResponse struct{}"
	}

	return strings.Join([]string{"ListBcsMetricResponse", string(data)}, " ")
}
