package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateDomainGroupInheritRoleRequest struct {
	DomainId string `json:"domain_id"`
	GroupId  string `json:"group_id"`
	RoleId   string `json:"role_id"`
}

func (o UpdateDomainGroupInheritRoleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDomainGroupInheritRoleRequest struct{}"
	}

	return strings.Join([]string{"UpdateDomainGroupInheritRoleRequest", string(data)}, " ")
}
