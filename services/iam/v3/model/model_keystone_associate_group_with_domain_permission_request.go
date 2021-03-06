package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneAssociateGroupWithDomainPermissionRequest struct {
	DomainId string `json:"domain_id"`
	GroupId  string `json:"group_id"`
	RoleId   string `json:"role_id"`
}

func (o KeystoneAssociateGroupWithDomainPermissionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneAssociateGroupWithDomainPermissionRequest struct{}"
	}

	return strings.Join([]string{"KeystoneAssociateGroupWithDomainPermissionRequest", string(data)}, " ")
}
