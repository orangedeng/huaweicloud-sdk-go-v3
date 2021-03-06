package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneListVersionsResponse struct {
	Versions       *Versions `json:"versions,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o KeystoneListVersionsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneListVersionsResponse struct{}"
	}

	return strings.Join([]string{"KeystoneListVersionsResponse", string(data)}, " ")
}
