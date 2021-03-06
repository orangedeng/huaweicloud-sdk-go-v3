package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type AssociateVaultPolicyRequest struct {
	VaultId string          `json:"vault_id"`
	Body    *VaultAssociate `json:"body,omitempty"`
}

func (o AssociateVaultPolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AssociateVaultPolicyRequest struct{}"
	}

	return strings.Join([]string{"AssociateVaultPolicyRequest", string(data)}, " ")
}
