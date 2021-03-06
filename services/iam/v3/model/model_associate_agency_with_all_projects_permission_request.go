package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type AssociateAgencyWithAllProjectsPermissionRequest struct {
	AgencyId string `json:"agency_id"`
	DomainId string `json:"domain_id"`
	RoleId   string `json:"role_id"`
}

func (o AssociateAgencyWithAllProjectsPermissionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AssociateAgencyWithAllProjectsPermissionRequest struct{}"
	}

	return strings.Join([]string{"AssociateAgencyWithAllProjectsPermissionRequest", string(data)}, " ")
}
