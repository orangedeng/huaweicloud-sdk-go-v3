package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateAlarmActionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAlarmActionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateAlarmActionResponse struct{}"
	}

	return strings.Join([]string{"UpdateAlarmActionResponse", string(data)}, " ")
}
