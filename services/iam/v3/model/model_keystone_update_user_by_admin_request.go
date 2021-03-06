package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneUpdateUserByAdminRequest struct {
	UserId string                                `json:"user_id"`
	Body   *KeystoneUpdateUserByAdminRequestBody `json:"body,omitempty"`
}

func (o KeystoneUpdateUserByAdminRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneUpdateUserByAdminRequest struct{}"
	}

	return strings.Join([]string{"KeystoneUpdateUserByAdminRequest", string(data)}, " ")
}
