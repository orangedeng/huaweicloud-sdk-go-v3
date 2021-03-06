package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateDomainConsoleAclPolicyRequest struct {
	DomainId string                                   `json:"domain_id"`
	Body     *UpdateDomainConsoleAclPolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateDomainConsoleAclPolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDomainConsoleAclPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateDomainConsoleAclPolicyRequest", string(data)}, " ")
}
