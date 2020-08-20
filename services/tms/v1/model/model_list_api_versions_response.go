/*
 * TMS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListApiVersionsResponse struct {
	// 版本列表
	Versions []VersionDetail `json:"versions,omitempty"`
}

func (o ListApiVersionsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListApiVersionsResponse", string(data)}, " ")
}
