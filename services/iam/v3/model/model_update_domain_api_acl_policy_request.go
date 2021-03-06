package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateDomainApiAclPolicyRequest struct {
	DomainId string                               `json:"domain_id"`
	Body     *UpdateDomainApiAclPolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateDomainApiAclPolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDomainApiAclPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateDomainApiAclPolicyRequest", string(data)}, " ")
}
