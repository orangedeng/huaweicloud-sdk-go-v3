package model

import (
	"encoding/json"

	"strings"
)

// 添加节点和添加组织，添加组织是需要提供pvc_name
type UpdateInstanceRequestBody struct {
	// 添加节点的组织列表
	NodeOrgs []NodeOrgs `json:"node_orgs"`
}

func (o UpdateInstanceRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateInstanceRequestBody", string(data)}, " ")
}
