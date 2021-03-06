package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowL7policyRequest struct {
	L7policyId string `json:"l7policy_id"`
}

func (o ShowL7policyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowL7policyRequest struct{}"
	}

	return strings.Join([]string{"ShowL7policyRequest", string(data)}, " ")
}
