package model

import (
	"encoding/json"

	"strings"
)

// 节点组织
type NodeOrgs struct {
	// 组织名称
	Name string `json:"name"`
	// 组织目标节点数
	NodeCount int32 `json:"node_count"`
	// pvc名称，添加组织时需要提供pvc_name
	PvcName *string `json:"pvc_name,omitempty"`
}

func (o NodeOrgs) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NodeOrgs struct{}"
	}

	return strings.Join([]string{"NodeOrgs", string(data)}, " ")
}
