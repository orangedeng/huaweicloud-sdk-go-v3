package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneRemoveUserFromGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o KeystoneRemoveUserFromGroupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneRemoveUserFromGroupResponse struct{}"
	}

	return strings.Join([]string{"KeystoneRemoveUserFromGroupResponse", string(data)}, " ")
}
