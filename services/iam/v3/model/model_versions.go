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

//
type Versions struct {
	// 版本的资源链接信息。
	Values []Version `json:"values"`
}

func (o Versions) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"Versions", string(data)}, " ")
}
