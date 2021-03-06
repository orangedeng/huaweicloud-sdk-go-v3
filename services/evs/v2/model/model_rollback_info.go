package model

import (
	"encoding/json"

	"strings"
)

type RollbackInfo struct {
	// 回滚的目标云硬盘UUID。
	VolumeId string `json:"volume_id"`
}

func (o RollbackInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RollbackInfo struct{}"
	}

	return strings.Join([]string{"RollbackInfo", string(data)}, " ")
}
