package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateAlarmResponse struct {
	// 告警规则的ID。
	AlarmId        *string `json:"alarm_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAlarmResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateAlarmResponse struct{}"
	}

	return strings.Join([]string{"CreateAlarmResponse", string(data)}, " ")
}
