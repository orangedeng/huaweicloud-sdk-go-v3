package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneDeleteIdentityProviderResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o KeystoneDeleteIdentityProviderResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneDeleteIdentityProviderResponse struct{}"
	}

	return strings.Join([]string{"KeystoneDeleteIdentityProviderResponse", string(data)}, " ")
}
