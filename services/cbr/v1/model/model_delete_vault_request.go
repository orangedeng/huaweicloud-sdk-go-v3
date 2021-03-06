package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteVaultRequest struct {
	VaultId string `json:"vault_id"`
}

func (o DeleteVaultRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteVaultRequest struct{}"
	}

	return strings.Join([]string{"DeleteVaultRequest", string(data)}, " ")
}
