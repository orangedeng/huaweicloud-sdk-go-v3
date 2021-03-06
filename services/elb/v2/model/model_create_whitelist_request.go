package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateWhitelistRequest struct {
	Body *CreateWhitelistRequestBody `json:"body,omitempty"`
}

func (o CreateWhitelistRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateWhitelistRequest struct{}"
	}

	return strings.Join([]string{"CreateWhitelistRequest", string(data)}, " ")
}
