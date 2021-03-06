package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteDomainGroupInheritedRoleRequest struct {
	DomainId string `json:"domain_id"`
	GroupId  string `json:"group_id"`
	RoleId   string `json:"role_id"`
}

func (o DeleteDomainGroupInheritedRoleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteDomainGroupInheritedRoleRequest struct{}"
	}

	return strings.Join([]string{"DeleteDomainGroupInheritedRoleRequest", string(data)}, " ")
}
