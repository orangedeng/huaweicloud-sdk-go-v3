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

// Request Object
type KeystoneListProjectPermissionsForGroupRequest struct {
	ProjectId string `json:"project_id"`
	GroupId   string `json:"group_id"`
}

func (o KeystoneListProjectPermissionsForGroupRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneListProjectPermissionsForGroupRequest", string(data)}, " ")
}
