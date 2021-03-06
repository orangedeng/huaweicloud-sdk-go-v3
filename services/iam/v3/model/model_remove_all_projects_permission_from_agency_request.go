package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RemoveAllProjectsPermissionFromAgencyRequest struct {
	AgencyId string `json:"agency_id"`
	DomainId string `json:"domain_id"`
	RoleId   string `json:"role_id"`
}

func (o RemoveAllProjectsPermissionFromAgencyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RemoveAllProjectsPermissionFromAgencyRequest struct{}"
	}

	return strings.Join([]string{"RemoveAllProjectsPermissionFromAgencyRequest", string(data)}, " ")
}
