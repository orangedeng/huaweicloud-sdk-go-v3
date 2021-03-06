package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteAlarmTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAlarmTemplateResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteAlarmTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteAlarmTemplateResponse", string(data)}, " ")
}
