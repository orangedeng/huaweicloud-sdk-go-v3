package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneCheckUserInGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o KeystoneCheckUserInGroupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneCheckUserInGroupResponse struct{}"
	}

	return strings.Join([]string{"KeystoneCheckUserInGroupResponse", string(data)}, " ")
}
