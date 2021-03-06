package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneDeleteUserResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o KeystoneDeleteUserResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneDeleteUserResponse struct{}"
	}

	return strings.Join([]string{"KeystoneDeleteUserResponse", string(data)}, " ")
}
