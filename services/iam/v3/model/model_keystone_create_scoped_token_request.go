package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneCreateScopedTokenRequest struct {
	Body *KeystoneCreateScopedTokenRequestBody `json:"body,omitempty"`
}

func (o KeystoneCreateScopedTokenRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneCreateScopedTokenRequest struct{}"
	}

	return strings.Join([]string{"KeystoneCreateScopedTokenRequest", string(data)}, " ")
}
