package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneCheckroleForGroupRequest struct {
	DomainId string `json:"domain_id"`
	GroupId  string `json:"group_id"`
	RoleId   string `json:"role_id"`
}

func (o KeystoneCheckroleForGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneCheckroleForGroupRequest struct{}"
	}

	return strings.Join([]string{"KeystoneCheckroleForGroupRequest", string(data)}, " ")
}
