package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneAssociateGroupWithProjectPermissionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o KeystoneAssociateGroupWithProjectPermissionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneAssociateGroupWithProjectPermissionResponse struct{}"
	}

	return strings.Join([]string{"KeystoneAssociateGroupWithProjectPermissionResponse", string(data)}, " ")
}
