package model

import (
	"encoding/json"

	"strings"
)

// 备份个数，键值对，包含：reserved（预留）、allocated（预留）、limit（最大）和in_use（已使用）。
type QuotaDetailBackups struct {
	// 已使用的数量。
	InUse int32 `json:"in_use"`
	// 最大的数量。
	Limit int32 `json:"limit"`
	// 预留属性。
	Reserved int32 `json:"reserved"`
	// 预留属性。
	Allocated int32 `json:"allocated"`
}

func (o QuotaDetailBackups) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QuotaDetailBackups struct{}"
	}

	return strings.Join([]string{"QuotaDetailBackups", string(data)}, " ")
}
