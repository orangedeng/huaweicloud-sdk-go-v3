package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneCreateUserTokenByPasswordAndMfaRequest struct {
	Nocatalog *string                                             `json:"nocatalog,omitempty"`
	Body      *KeystoneCreateUserTokenByPasswordAndMfaRequestBody `json:"body,omitempty"`
}

func (o KeystoneCreateUserTokenByPasswordAndMfaRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneCreateUserTokenByPasswordAndMfaRequest struct{}"
	}

	return strings.Join([]string{"KeystoneCreateUserTokenByPasswordAndMfaRequest", string(data)}, " ")
}
