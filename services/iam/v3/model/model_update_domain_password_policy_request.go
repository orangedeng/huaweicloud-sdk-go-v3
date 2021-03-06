package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateDomainPasswordPolicyRequest struct {
	DomainId string                                 `json:"domain_id"`
	Body     *UpdateDomainPasswordPolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateDomainPasswordPolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDomainPasswordPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateDomainPasswordPolicyRequest", string(data)}, " ")
}
