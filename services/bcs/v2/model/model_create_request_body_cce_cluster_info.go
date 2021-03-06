package model

import (
	"encoding/json"

	"strings"
)

// 使用已有CCE集群信息
type CreateRequestBodyCceClusterInfo struct {
	// 集群ID
	ClusterId *string `json:"cluster_id,omitempty"`
	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`
}

func (o CreateRequestBodyCceClusterInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateRequestBodyCceClusterInfo struct{}"
	}

	return strings.Join([]string{"CreateRequestBodyCceClusterInfo", string(data)}, " ")
}
