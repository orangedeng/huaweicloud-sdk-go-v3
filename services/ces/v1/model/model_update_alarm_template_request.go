package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateAlarmTemplateRequest struct {
	ContentType string                          `json:"Content-Type"`
	TemplateId  string                          `json:"template_id"`
	Body        *UpdateAlarmTemplateRequestBody `json:"body,omitempty"`
}

func (o UpdateAlarmTemplateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateAlarmTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateAlarmTemplateRequest", string(data)}, " ")
}
