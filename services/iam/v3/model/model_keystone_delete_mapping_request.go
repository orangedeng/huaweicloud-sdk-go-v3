package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneDeleteMappingRequest struct {
	Id string `json:"id"`
}

func (o KeystoneDeleteMappingRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneDeleteMappingRequest struct{}"
	}

	return strings.Join([]string{"KeystoneDeleteMappingRequest", string(data)}, " ")
}
