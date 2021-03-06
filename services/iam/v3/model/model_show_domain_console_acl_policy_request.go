package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowDomainConsoleAclPolicyRequest struct {
	DomainId string `json:"domain_id"`
}

func (o ShowDomainConsoleAclPolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDomainConsoleAclPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainConsoleAclPolicyRequest", string(data)}, " ")
}
