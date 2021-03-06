package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type AssociateAgencyWithProjectPermissionRequest struct {
	ProjectId string `json:"project_id"`
	AgencyId  string `json:"agency_id"`
	RoleId    string `json:"role_id"`
}

func (o AssociateAgencyWithProjectPermissionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AssociateAgencyWithProjectPermissionRequest struct{}"
	}

	return strings.Join([]string{"AssociateAgencyWithProjectPermissionRequest", string(data)}, " ")
}
