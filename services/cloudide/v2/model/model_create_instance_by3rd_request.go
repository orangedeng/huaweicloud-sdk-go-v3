/*
 * cloudide
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
type CreateInstanceBy3rdRequest struct {
	InstanceLabel *string            `json:"instance_label,omitempty"`
	Body          *InstanceEdgeParam `json:"body,omitempty"`
}

func (o CreateInstanceBy3rdRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateInstanceBy3rdRequest", string(data)}, " ")
}
