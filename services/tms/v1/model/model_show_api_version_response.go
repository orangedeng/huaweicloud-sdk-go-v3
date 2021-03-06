package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowApiVersionResponse struct {
	Version        *VersionDetail `json:"version,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowApiVersionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowApiVersionResponse struct{}"
	}

	return strings.Join([]string{"ShowApiVersionResponse", string(data)}, " ")
}
