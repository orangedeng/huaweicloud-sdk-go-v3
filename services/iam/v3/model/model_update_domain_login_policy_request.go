package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateDomainLoginPolicyRequest struct {
	DomainId string                              `json:"domain_id"`
	Body     *UpdateDomainLoginPolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateDomainLoginPolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDomainLoginPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateDomainLoginPolicyRequest", string(data)}, " ")
}
