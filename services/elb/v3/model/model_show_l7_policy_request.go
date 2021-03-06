package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowL7PolicyRequest struct {
	L7policyId string `json:"l7policy_id"`
}

func (o ShowL7PolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowL7PolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowL7PolicyRequest", string(data)}, " ")
}
