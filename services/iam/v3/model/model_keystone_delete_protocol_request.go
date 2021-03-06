package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneDeleteProtocolRequest struct {
	IdpId      string `json:"idp_id"`
	ProtocolId string `json:"protocol_id"`
}

func (o KeystoneDeleteProtocolRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneDeleteProtocolRequest struct{}"
	}

	return strings.Join([]string{"KeystoneDeleteProtocolRequest", string(data)}, " ")
}
