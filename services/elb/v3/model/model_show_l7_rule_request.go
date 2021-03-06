package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowL7RuleRequest struct {
	L7policyId string `json:"l7policy_id"`
	L7ruleId   string `json:"l7rule_id"`
}

func (o ShowL7RuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowL7RuleRequest struct{}"
	}

	return strings.Join([]string{"ShowL7RuleRequest", string(data)}, " ")
}
