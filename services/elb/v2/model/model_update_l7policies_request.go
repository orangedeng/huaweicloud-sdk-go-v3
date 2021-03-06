package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateL7policiesRequest struct {
	L7policyId string                       `json:"l7policy_id"`
	Body       *UpdateL7policiesRequestBody `json:"body,omitempty"`
}

func (o UpdateL7policiesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateL7policiesRequest struct{}"
	}

	return strings.Join([]string{"UpdateL7policiesRequest", string(data)}, " ")
}
