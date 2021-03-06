package model

import (
	"encoding/json"

	"strings"
)

// quota响应
type Quota struct {
	// 查询的配额信息。
	Resources []QuotaInfo `json:"resources"`
}

func (o Quota) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Quota struct{}"
	}

	return strings.Join([]string{"Quota", string(data)}, " ")
}
