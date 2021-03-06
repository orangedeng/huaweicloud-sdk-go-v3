package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneUpdateUserByAdminResponse struct {
	User           *KeystoneUpdateUserByAdminResult `json:"user,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o KeystoneUpdateUserByAdminResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneUpdateUserByAdminResponse struct{}"
	}

	return strings.Join([]string{"KeystoneUpdateUserByAdminResponse", string(data)}, " ")
}
