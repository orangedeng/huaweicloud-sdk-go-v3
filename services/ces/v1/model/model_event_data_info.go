package model

import (
	"encoding/json"

	"strings"
)

//
type EventDataInfo struct {
	// 事件类型，例如instance_host_info。
	Type string `json:"type"`
	// 事件上报时间。
	Timestamp int64 `json:"timestamp"`
	// 主机配置信息。
	Value string `json:"value"`
}

func (o EventDataInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "EventDataInfo struct{}"
	}

	return strings.Join([]string{"EventDataInfo", string(data)}, " ")
}
