package model

import (
	"encoding/json"

	"strings"
)

// 配额详细信息。
type QuotaDetail struct {
	// 已使用的数量。
	InUse int32 `json:"in_use"`
	// 最大的数量。
	Limit int32 `json:"limit"`
	// 预留属性。
	Reserved int32 `json:"reserved"`
	// 预留属性。
	Allocated int32 `json:"allocated"`
}

func (o QuotaDetail) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QuotaDetail struct{}"
	}

	return strings.Join([]string{"QuotaDetail", string(data)}, " ")
}
