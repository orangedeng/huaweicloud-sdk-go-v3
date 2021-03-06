package model

import (
	"encoding/json"

	"strings"
)

// 极速文件存储卷信息
type CreateRequestBodyTurboInfo struct {
	// 共享方式
	ShareType *string `json:"share_type,omitempty"`
	// 类型
	Type *string `json:"type,omitempty"`
	// 可用区
	AvailableZone *string `json:"available_zone,omitempty"`
	// 规格
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`
}

func (o CreateRequestBodyTurboInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateRequestBodyTurboInfo struct{}"
	}

	return strings.Join([]string{"CreateRequestBodyTurboInfo", string(data)}, " ")
}
