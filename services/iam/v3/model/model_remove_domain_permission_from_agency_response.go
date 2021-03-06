package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RemoveDomainPermissionFromAgencyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RemoveDomainPermissionFromAgencyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RemoveDomainPermissionFromAgencyResponse struct{}"
	}

	return strings.Join([]string{"RemoveDomainPermissionFromAgencyResponse", string(data)}, " ")
}
