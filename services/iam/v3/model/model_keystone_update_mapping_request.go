package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneUpdateMappingRequest struct {
	Id   string                            `json:"id"`
	Body *KeystoneUpdateMappingRequestBody `json:"body,omitempty"`
}

func (o KeystoneUpdateMappingRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneUpdateMappingRequest struct{}"
	}

	return strings.Join([]string{"KeystoneUpdateMappingRequest", string(data)}, " ")
}
