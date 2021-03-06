package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneCheckroleForGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o KeystoneCheckroleForGroupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneCheckroleForGroupResponse struct{}"
	}

	return strings.Join([]string{"KeystoneCheckroleForGroupResponse", string(data)}, " ")
}
