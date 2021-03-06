package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneListAllProjectPermissionsForGroupRequest struct {
	DomainId string `json:"domain_id"`
	GroupId  string `json:"group_id"`
}

func (o KeystoneListAllProjectPermissionsForGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneListAllProjectPermissionsForGroupRequest struct{}"
	}

	return strings.Join([]string{"KeystoneListAllProjectPermissionsForGroupRequest", string(data)}, " ")
}
