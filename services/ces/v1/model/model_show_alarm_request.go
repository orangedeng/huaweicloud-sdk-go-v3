package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowAlarmRequest struct {
	ContentType string `json:"Content-Type"`
	AlarmId     string `json:"alarm_id"`
}

func (o ShowAlarmRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowAlarmRequest struct{}"
	}

	return strings.Join([]string{"ShowAlarmRequest", string(data)}, " ")
}
