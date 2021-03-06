package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateTemporaryAccessKeyByTokenResponse struct {
	Credential     *Credential `json:"credential,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o CreateTemporaryAccessKeyByTokenResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateTemporaryAccessKeyByTokenResponse struct{}"
	}

	return strings.Join([]string{"CreateTemporaryAccessKeyByTokenResponse", string(data)}, " ")
}
