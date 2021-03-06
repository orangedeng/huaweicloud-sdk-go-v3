package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneCheckProjectPermissionForGroupRequest struct {
	ProjectId string `json:"project_id"`
	GroupId   string `json:"group_id"`
	RoleId    string `json:"role_id"`
}

func (o KeystoneCheckProjectPermissionForGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneCheckProjectPermissionForGroupRequest struct{}"
	}

	return strings.Join([]string{"KeystoneCheckProjectPermissionForGroupRequest", string(data)}, " ")
}
