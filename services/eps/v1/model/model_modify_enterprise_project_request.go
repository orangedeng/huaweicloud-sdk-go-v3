package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ModifyEnterpriseProjectRequest struct {
	EnterpriseProjectId string             `json:"enterprise_project_id"`
	Body                *EnterpriseProject `json:"body,omitempty"`
}

func (o ModifyEnterpriseProjectRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ModifyEnterpriseProjectRequest struct{}"
	}

	return strings.Join([]string{"ModifyEnterpriseProjectRequest", string(data)}, " ")
}
