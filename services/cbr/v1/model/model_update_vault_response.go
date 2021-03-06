package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateVaultResponse struct {
	Vault          *Vault `json:"vault,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdateVaultResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateVaultResponse struct{}"
	}

	return strings.Join([]string{"UpdateVaultResponse", string(data)}, " ")
}
