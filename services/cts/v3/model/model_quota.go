package model

import (
	"encoding/json"

	"strings"
)

type Quota struct {
	// quota资源类型。
	Type *string `json:"type,omitempty"`
	// 已使用的资源个数。
	Used *int64 `json:"used,omitempty"`
	// 总资源个数。
	Quota *int64 `json:"quota,omitempty"`
}

func (o Quota) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Quota struct{}"
	}

	return strings.Join([]string{"Quota", string(data)}, " ")
}
