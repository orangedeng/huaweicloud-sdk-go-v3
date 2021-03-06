package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchAddPeersToChannelRequest struct {
	BlockchainId string                             `json:"blockchain_id"`
	Body         *BatchAddPeersToChannelRequestBody `json:"body,omitempty"`
}

func (o BatchAddPeersToChannelRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchAddPeersToChannelRequest struct{}"
	}

	return strings.Join([]string{"BatchAddPeersToChannelRequest", string(data)}, " ")
}
