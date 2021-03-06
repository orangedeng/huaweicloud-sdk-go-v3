package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneDeleteIdentityProviderRequest struct {
	Id string `json:"id"`
}

func (o KeystoneDeleteIdentityProviderRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneDeleteIdentityProviderRequest struct{}"
	}

	return strings.Join([]string{"KeystoneDeleteIdentityProviderRequest", string(data)}, " ")
}
