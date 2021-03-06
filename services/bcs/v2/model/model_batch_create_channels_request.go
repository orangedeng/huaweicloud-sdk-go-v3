package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchCreateChannelsRequest struct {
	BlockchainId string                          `json:"blockchain_id"`
	Body         *BatchCreateChannelsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateChannelsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchCreateChannelsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateChannelsRequest", string(data)}, " ")
}
