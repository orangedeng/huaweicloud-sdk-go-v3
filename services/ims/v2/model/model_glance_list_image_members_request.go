/*
 * IMS
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
type GlanceListImageMembersRequest struct {
	ImageId string `json:"image_id"`
}

func (o GlanceListImageMembersRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"GlanceListImageMembersRequest", string(data)}, " ")
}
