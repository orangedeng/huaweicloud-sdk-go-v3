package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowVaultResourceInstancesRequest struct {
	Body *VaultResourceInstancesReq `json:"body,omitempty"`
}

func (o ShowVaultResourceInstancesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowVaultResourceInstancesRequest struct{}"
	}

	return strings.Join([]string{"ShowVaultResourceInstancesRequest", string(data)}, " ")
}
