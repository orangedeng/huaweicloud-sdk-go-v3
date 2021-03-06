package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DisableEnterpriseProjectResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DisableEnterpriseProjectResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DisableEnterpriseProjectResponse struct{}"
	}

	return strings.Join([]string{"DisableEnterpriseProjectResponse", string(data)}, " ")
}
