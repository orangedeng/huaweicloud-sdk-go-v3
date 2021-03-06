package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneListDomainPermissionsForGroupResponse struct {
	Links *Links `json:"links,omitempty"`
	// 权限信息列表。
	Roles          *[]RoleResult `json:"roles,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o KeystoneListDomainPermissionsForGroupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneListDomainPermissionsForGroupResponse struct{}"
	}

	return strings.Join([]string{"KeystoneListDomainPermissionsForGroupResponse", string(data)}, " ")
}
