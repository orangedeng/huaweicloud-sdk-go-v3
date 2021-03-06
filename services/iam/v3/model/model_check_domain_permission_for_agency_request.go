package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CheckDomainPermissionForAgencyRequest struct {
	DomainId string `json:"domain_id"`
	AgencyId string `json:"agency_id"`
	RoleId   string `json:"role_id"`
}

func (o CheckDomainPermissionForAgencyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CheckDomainPermissionForAgencyRequest struct{}"
	}

	return strings.Join([]string{"CheckDomainPermissionForAgencyRequest", string(data)}, " ")
}
