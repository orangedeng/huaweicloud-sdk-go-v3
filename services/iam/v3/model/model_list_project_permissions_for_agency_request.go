package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListProjectPermissionsForAgencyRequest struct {
	ProjectId string `json:"project_id"`
	AgencyId  string `json:"agency_id"`
}

func (o ListProjectPermissionsForAgencyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListProjectPermissionsForAgencyRequest struct{}"
	}

	return strings.Join([]string{"ListProjectPermissionsForAgencyRequest", string(data)}, " ")
}
