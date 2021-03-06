package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowDomainPasswordPolicyRequest struct {
	DomainId string `json:"domain_id"`
}

func (o ShowDomainPasswordPolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDomainPasswordPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainPasswordPolicyRequest", string(data)}, " ")
}
