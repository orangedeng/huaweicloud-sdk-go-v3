/*
 * IAM
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
type ListDomainPermissionsForAgencyResponse struct {
	// 权限信息列表。
	Roles []RoleResult `json:"roles,omitempty"`
}

func (o ListDomainPermissionsForAgencyResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListDomainPermissionsForAgencyResponse", string(data)}, " ")
}
