/*
 * Cbr
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateVaultRequest struct {
	VaultId string          `json:"vault_id"`
	Body    *VaultUpdateReq `json:"body,omitempty"`
}

func (o UpdateVaultRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateVaultRequest", string(data)}, " ")
}
