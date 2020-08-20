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

type VaultRemoveResourceReq struct {
	// 要移除的资源ID列表
	ResourceIds []string `json:"resource_ids"`
}

func (o VaultRemoveResourceReq) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"VaultRemoveResourceReq", string(data)}, " ")
}
