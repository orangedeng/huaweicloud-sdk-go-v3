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

type OpExtendInfoRemoveResources struct {
	// 删除失败的资源ID列表
	FailCount *int32 `json:"fail_count,omitempty"`
	// 删除的备份数量
	TotalCount *int32 `json:"total_count,omitempty"`
	//
	Resources []Resource `json:"resources,omitempty"`
}

func (o OpExtendInfoRemoveResources) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"OpExtendInfoRemoveResources", string(data)}, " ")
}
