package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneRemoveDomainPermissionFromGroupRequest struct {
	DomainId string `json:"domain_id"`
	GroupId  string `json:"group_id"`
	RoleId   string `json:"role_id"`
}

func (o KeystoneRemoveDomainPermissionFromGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneRemoveDomainPermissionFromGroupRequest struct{}"
	}

	return strings.Join([]string{"KeystoneRemoveDomainPermissionFromGroupRequest", string(data)}, " ")
}
