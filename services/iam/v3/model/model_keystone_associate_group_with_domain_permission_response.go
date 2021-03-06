package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneAssociateGroupWithDomainPermissionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o KeystoneAssociateGroupWithDomainPermissionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneAssociateGroupWithDomainPermissionResponse struct{}"
	}

	return strings.Join([]string{"KeystoneAssociateGroupWithDomainPermissionResponse", string(data)}, " ")
}
