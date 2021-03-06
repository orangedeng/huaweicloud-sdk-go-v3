package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneListProtocolsRequest struct {
	IdpId string `json:"idp_id"`
}

func (o KeystoneListProtocolsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneListProtocolsRequest struct{}"
	}

	return strings.Join([]string{"KeystoneListProtocolsRequest", string(data)}, " ")
}
