package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreatePolicyRequest struct {
	Body *PolicyCreateReq `json:"body,omitempty"`
}

func (o CreatePolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreatePolicyRequest struct{}"
	}

	return strings.Join([]string{"CreatePolicyRequest", string(data)}, " ")
}
