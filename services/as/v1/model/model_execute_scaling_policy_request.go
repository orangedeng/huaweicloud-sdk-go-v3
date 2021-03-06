package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ExecuteScalingPolicyRequest struct {
	ScalingPolicyId string                           `json:"scaling_policy_id"`
	Body            *ExecuteScalingPolicyRequestBody `json:"body,omitempty"`
}

func (o ExecuteScalingPolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExecuteScalingPolicyRequest struct{}"
	}

	return strings.Join([]string{"ExecuteScalingPolicyRequest", string(data)}, " ")
}
