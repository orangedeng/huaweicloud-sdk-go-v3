package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateNewBlockchainResponse struct {
	// 服务实例ID
	BlockchainId *string `json:"blockchain_id,omitempty"`
	// 服务实例名
	BlockchainName *string `json:"blockchain_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateNewBlockchainResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateNewBlockchainResponse struct{}"
	}

	return strings.Join([]string{"CreateNewBlockchainResponse", string(data)}, " ")
}
