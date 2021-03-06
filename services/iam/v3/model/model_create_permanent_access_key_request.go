package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreatePermanentAccessKeyRequest struct {
	Body *CreatePermanentAccessKeyRequestBody `json:"body,omitempty"`
}

func (o CreatePermanentAccessKeyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreatePermanentAccessKeyRequest struct{}"
	}

	return strings.Join([]string{"CreatePermanentAccessKeyRequest", string(data)}, " ")
}
