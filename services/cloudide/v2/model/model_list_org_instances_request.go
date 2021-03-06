package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListOrgInstancesRequest struct {
	IsTemporary *bool   `json:"is_temporary,omitempty"`
	Limit       *int64  `json:"limit,omitempty"`
	Offset      *int64  `json:"offset,omitempty"`
	OrgId       string  `json:"org_id"`
	Search      *string `json:"search,omitempty"`
}

func (o ListOrgInstancesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListOrgInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListOrgInstancesRequest", string(data)}, " ")
}
