package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type AssociateVaultPolicyResponse struct {
	AssociatePolicy *VaultPolicyResp `json:"associate_policy,omitempty"`
	HttpStatusCode  int              `json:"-"`
}

func (o AssociateVaultPolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AssociateVaultPolicyResponse struct{}"
	}

	return strings.Join([]string{"AssociateVaultPolicyResponse", string(data)}, " ")
}
