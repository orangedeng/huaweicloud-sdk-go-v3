package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowMetadataRequest struct {
	IdpId      string `json:"idp_id"`
	ProtocolId string `json:"protocol_id"`
}

func (o ShowMetadataRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowMetadataRequest struct{}"
	}

	return strings.Join([]string{"ShowMetadataRequest", string(data)}, " ")
}
