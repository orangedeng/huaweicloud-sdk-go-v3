package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateEventsRequest struct {
	ContentType string `json:"Content-Type"`
	// 上报自定义事件。请求参数。
	Body *[]EventItem `json:"body,omitempty"`
}

func (o CreateEventsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateEventsRequest struct{}"
	}

	return strings.Join([]string{"CreateEventsRequest", string(data)}, " ")
}
