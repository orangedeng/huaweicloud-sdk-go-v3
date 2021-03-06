package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneListGroupsRequest struct {
	DomainId *string `json:"domain_id,omitempty"`
	Name     *string `json:"name,omitempty"`
}

func (o KeystoneListGroupsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneListGroupsRequest struct{}"
	}

	return strings.Join([]string{"KeystoneListGroupsRequest", string(data)}, " ")
}
