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

type VaultAddResourceReq struct {
	// 资源列表
	Resources []ResourceCreate `json:"resources"`
}

func (o VaultAddResourceReq) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"VaultAddResourceReq", string(data)}, " ")
}
