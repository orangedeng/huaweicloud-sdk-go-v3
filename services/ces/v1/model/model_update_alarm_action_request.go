package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateAlarmActionRequest struct {
	ContentType string                `json:"Content-Type"`
	AlarmId     string                `json:"alarm_id"`
	Body        *ModifyAlarmActionReq `json:"body,omitempty"`
}

func (o UpdateAlarmActionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateAlarmActionRequest struct{}"
	}

	return strings.Join([]string{"UpdateAlarmActionRequest", string(data)}, " ")
}
