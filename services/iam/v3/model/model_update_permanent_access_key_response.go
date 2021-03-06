package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdatePermanentAccessKeyResponse struct {
	Credential     *UpdateCredentialResult `json:"credential,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o UpdatePermanentAccessKeyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePermanentAccessKeyResponse struct{}"
	}

	return strings.Join([]string{"UpdatePermanentAccessKeyResponse", string(data)}, " ")
}
