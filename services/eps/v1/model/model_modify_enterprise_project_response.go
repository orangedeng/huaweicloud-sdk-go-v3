package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ModifyEnterpriseProjectResponse struct {
	EnterpriseProject *EpDetail `json:"enterprise_project,omitempty"`
	HttpStatusCode    int       `json:"-"`
}

func (o ModifyEnterpriseProjectResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ModifyEnterpriseProjectResponse struct{}"
	}

	return strings.Join([]string{"ModifyEnterpriseProjectResponse", string(data)}, " ")
}
