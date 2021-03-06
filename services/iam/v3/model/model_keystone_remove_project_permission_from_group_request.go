package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneRemoveProjectPermissionFromGroupRequest struct {
	ProjectId string `json:"project_id"`
	GroupId   string `json:"group_id"`
	RoleId    string `json:"role_id"`
}

func (o KeystoneRemoveProjectPermissionFromGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneRemoveProjectPermissionFromGroupRequest struct{}"
	}

	return strings.Join([]string{"KeystoneRemoveProjectPermissionFromGroupRequest", string(data)}, " ")
}
