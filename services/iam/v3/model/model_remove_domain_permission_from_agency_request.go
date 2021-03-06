package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RemoveDomainPermissionFromAgencyRequest struct {
	DomainId string `json:"domain_id"`
	AgencyId string `json:"agency_id"`
	RoleId   string `json:"role_id"`
}

func (o RemoveDomainPermissionFromAgencyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RemoveDomainPermissionFromAgencyRequest struct{}"
	}

	return strings.Join([]string{"RemoveDomainPermissionFromAgencyRequest", string(data)}, " ")
}
