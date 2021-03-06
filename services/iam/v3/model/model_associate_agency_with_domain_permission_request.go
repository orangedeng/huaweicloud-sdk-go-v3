package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type AssociateAgencyWithDomainPermissionRequest struct {
	DomainId string `json:"domain_id"`
	AgencyId string `json:"agency_id"`
	RoleId   string `json:"role_id"`
}

func (o AssociateAgencyWithDomainPermissionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AssociateAgencyWithDomainPermissionRequest struct{}"
	}

	return strings.Join([]string{"AssociateAgencyWithDomainPermissionRequest", string(data)}, " ")
}
