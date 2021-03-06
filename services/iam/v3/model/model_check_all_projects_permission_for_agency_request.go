package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CheckAllProjectsPermissionForAgencyRequest struct {
	AgencyId string `json:"agency_id"`
	DomainId string `json:"domain_id"`
	RoleId   string `json:"role_id"`
}

func (o CheckAllProjectsPermissionForAgencyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CheckAllProjectsPermissionForAgencyRequest struct{}"
	}

	return strings.Join([]string{"CheckAllProjectsPermissionForAgencyRequest", string(data)}, " ")
}
