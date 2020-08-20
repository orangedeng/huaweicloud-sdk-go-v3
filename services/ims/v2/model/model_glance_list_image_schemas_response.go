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

// Response Object
type GlanceListImageSchemasResponse struct {
	// 视图名称。
	Name *string `json:"name,omitempty"`
	// 镜像属性说明，主要是对基础属性的说明，包含每个属性的取值类型、用途等。
	Properties *interface{} `json:"properties,omitempty"`
	// 视图链接。
	Links []Links `json:"links,omitempty"`
}

func (o GlanceListImageSchemasResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"GlanceListImageSchemasResponse", string(data)}, " ")
}
