package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowEnterpriseProjectRequest struct {
	EnterpriseProjectId string `json:"enterprise_project_id"`
}

func (o ShowEnterpriseProjectRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowEnterpriseProjectRequest struct{}"
	}

	return strings.Join([]string{"ShowEnterpriseProjectRequest", string(data)}, " ")
}
