package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CheckProjectPermissionForAgencyRequest struct {
	ProjectId string `json:"project_id"`
	AgencyId  string `json:"agency_id"`
	RoleId    string `json:"role_id"`
}

func (o CheckProjectPermissionForAgencyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CheckProjectPermissionForAgencyRequest struct{}"
	}

	return strings.Join([]string{"CheckProjectPermissionForAgencyRequest", string(data)}, " ")
}
