package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteAlarmTemplateRequest struct {
	ContentType string `json:"Content-Type"`
	TemplateId  string `json:"template_id"`
}

func (o DeleteAlarmTemplateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteAlarmTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteAlarmTemplateRequest", string(data)}, " ")
}
