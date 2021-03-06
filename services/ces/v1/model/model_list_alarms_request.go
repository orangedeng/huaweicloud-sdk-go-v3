package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListAlarmsRequest struct {
	ContentType string  `json:"Content-Type"`
	Limit       *int32  `json:"limit,omitempty"`
	Order       *string `json:"order,omitempty"`
	Start       *string `json:"start,omitempty"`
}

func (o ListAlarmsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAlarmsRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmsRequest", string(data)}, " ")
}
