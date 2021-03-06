package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneDeleteUserRequest struct {
	UserId string `json:"user_id"`
}

func (o KeystoneDeleteUserRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneDeleteUserRequest struct{}"
	}

	return strings.Join([]string{"KeystoneDeleteUserRequest", string(data)}, " ")
}
