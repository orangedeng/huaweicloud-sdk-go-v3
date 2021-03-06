package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneUpdateUserPasswordRequest struct {
	UserId string                                 `json:"user_id"`
	Body   *KeystoneUpdateUserPasswordRequestBody `json:"body,omitempty"`
}

func (o KeystoneUpdateUserPasswordRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneUpdateUserPasswordRequest struct{}"
	}

	return strings.Join([]string{"KeystoneUpdateUserPasswordRequest", string(data)}, " ")
}
