package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowBlockchainStatusRequest struct {
	BlockchainId string `json:"blockchain_id"`
}

func (o ShowBlockchainStatusRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowBlockchainStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowBlockchainStatusRequest", string(data)}, " ")
}
