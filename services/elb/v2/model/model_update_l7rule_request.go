package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateL7ruleRequest struct {
	L7policyId string                   `json:"l7policy_id"`
	L7ruleId   string                   `json:"l7rule_id"`
	Body       *UpdateL7ruleRequestBody `json:"body,omitempty"`
}

func (o UpdateL7ruleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateL7ruleRequest struct{}"
	}

	return strings.Join([]string{"UpdateL7ruleRequest", string(data)}, " ")
}
