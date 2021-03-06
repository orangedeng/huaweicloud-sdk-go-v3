package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneUpdateMappingResponse struct {
	Mapping        *MappingResult `json:"mapping,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o KeystoneUpdateMappingResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneUpdateMappingResponse struct{}"
	}

	return strings.Join([]string{"KeystoneUpdateMappingResponse", string(data)}, " ")
}
