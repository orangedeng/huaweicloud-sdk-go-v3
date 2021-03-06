package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateAlarmTemplateRequest struct {
	ContentType string                          `json:"Content-Type"`
	Body        *CreateAlarmTemplateRequestBody `json:"body,omitempty"`
}

func (o CreateAlarmTemplateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateAlarmTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateAlarmTemplateRequest", string(data)}, " ")
}
