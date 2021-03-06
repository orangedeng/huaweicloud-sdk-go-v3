package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneUpdateProtocolRequest struct {
	IdpId      string                             `json:"idp_id"`
	ProtocolId string                             `json:"protocol_id"`
	Body       *KeystoneUpdateProtocolRequestBody `json:"body,omitempty"`
}

func (o KeystoneUpdateProtocolRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneUpdateProtocolRequest struct{}"
	}

	return strings.Join([]string{"KeystoneUpdateProtocolRequest", string(data)}, " ")
}
