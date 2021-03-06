package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneListPermissionsRequest struct {
	Name     *string `json:"name,omitempty"`
	DomainId *string `json:"domain_id,omitempty"`
	Page     *int32  `json:"page,omitempty"`
	PerPage  *int32  `json:"per_page,omitempty"`
}

func (o KeystoneListPermissionsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneListPermissionsRequest struct{}"
	}

	return strings.Join([]string{"KeystoneListPermissionsRequest", string(data)}, " ")
}
