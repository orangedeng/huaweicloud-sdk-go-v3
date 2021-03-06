package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateAgencyCustomPolicyRequest struct {
	RoleId string                               `json:"role_id"`
	Body   *UpdateAgencyCustomPolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateAgencyCustomPolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateAgencyCustomPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateAgencyCustomPolicyRequest", string(data)}, " ")
}
