package model

import (
	"encoding/json"

	"strings"
)

// 邀请方信息，被邀请方创建服务实例时需要填写此内容
type CreateRequestBodyInvitorInfos struct {
	// 邀请方租户ID
	TenantId *string `json:"tenant_id,omitempty"`
	// 邀请方项目ID
	ProjectId *string `json:"project_id,omitempty"`
	// 邀请方服务实例ID
	BlockchainId *string `json:"blockchain_id,omitempty"`
}

func (o CreateRequestBodyInvitorInfos) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateRequestBodyInvitorInfos struct{}"
	}

	return strings.Join([]string{"CreateRequestBodyInvitorInfos", string(data)}, " ")
}
