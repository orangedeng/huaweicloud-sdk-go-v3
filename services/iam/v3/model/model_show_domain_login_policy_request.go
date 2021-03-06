package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowDomainLoginPolicyRequest struct {
	DomainId string `json:"domain_id"`
}

func (o ShowDomainLoginPolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDomainLoginPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainLoginPolicyRequest", string(data)}, " ")
}
