package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateMetadataRequest struct {
	IdpId      string                     `json:"idp_id"`
	ProtocolId string                     `json:"protocol_id"`
	Body       *CreateMetadataRequestBody `json:"body,omitempty"`
}

func (o CreateMetadataRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateMetadataRequest struct{}"
	}

	return strings.Join([]string{"CreateMetadataRequest", string(data)}, " ")
}
