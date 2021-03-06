package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListCustomPoliciesRequest struct {
	Page    *int32 `json:"page,omitempty"`
	PerPage *int32 `json:"per_page,omitempty"`
}

func (o ListCustomPoliciesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListCustomPoliciesRequest struct{}"
	}

	return strings.Join([]string{"ListCustomPoliciesRequest", string(data)}, " ")
}
