package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneListServicesRequest struct {
	Type *string `json:"type,omitempty"`
}

func (o KeystoneListServicesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneListServicesRequest struct{}"
	}

	return strings.Join([]string{"KeystoneListServicesRequest", string(data)}, " ")
}
