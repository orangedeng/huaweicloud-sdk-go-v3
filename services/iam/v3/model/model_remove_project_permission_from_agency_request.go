package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RemoveProjectPermissionFromAgencyRequest struct {
	ProjectId string `json:"project_id"`
	AgencyId  string `json:"agency_id"`
	RoleId    string `json:"role_id"`
}

func (o RemoveProjectPermissionFromAgencyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RemoveProjectPermissionFromAgencyRequest struct{}"
	}

	return strings.Join([]string{"RemoveProjectPermissionFromAgencyRequest", string(data)}, " ")
}
