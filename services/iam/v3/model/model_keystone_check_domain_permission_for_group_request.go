package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneCheckDomainPermissionForGroupRequest struct {
	DomainId string `json:"domain_id"`
	GroupId  string `json:"group_id"`
	RoleId   string `json:"role_id"`
}

func (o KeystoneCheckDomainPermissionForGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneCheckDomainPermissionForGroupRequest struct{}"
	}

	return strings.Join([]string{"KeystoneCheckDomainPermissionForGroupRequest", string(data)}, " ")
}
