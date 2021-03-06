package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneListUsersRequest struct {
	DomainId          *string `json:"domain_id,omitempty"`
	Enabled           *bool   `json:"enabled,omitempty"`
	Name              *string `json:"name,omitempty"`
	PasswordExpiresAt *string `json:"password_expires_at,omitempty"`
}

func (o KeystoneListUsersRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneListUsersRequest struct{}"
	}

	return strings.Join([]string{"KeystoneListUsersRequest", string(data)}, " ")
}
