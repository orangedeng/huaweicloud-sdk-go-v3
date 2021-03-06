package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateEnterpriseProjectResponse struct {
	EnterpriseProject *EpDetail `json:"enterprise_project,omitempty"`
	HttpStatusCode    int       `json:"-"`
}

func (o CreateEnterpriseProjectResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateEnterpriseProjectResponse struct{}"
	}

	return strings.Join([]string{"CreateEnterpriseProjectResponse", string(data)}, " ")
}
