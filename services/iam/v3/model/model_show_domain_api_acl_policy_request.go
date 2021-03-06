package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowDomainApiAclPolicyRequest struct {
	DomainId string `json:"domain_id"`
}

func (o ShowDomainApiAclPolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDomainApiAclPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainApiAclPolicyRequest", string(data)}, " ")
}
