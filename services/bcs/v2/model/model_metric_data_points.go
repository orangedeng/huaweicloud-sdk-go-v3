package model

import (
	"encoding/json"

	"strings"
)

// 查询结果要点
type MetricDataPoints struct {
	// 时间戳。
	Timestamp *int64 `json:"timestamp,omitempty"`
	// 指标单位。
	Unit *string `json:"unit,omitempty"`
	// 统计方式。
	Statistics *[]StatisticValue `json:"statistics,omitempty"`
}

func (o MetricDataPoints) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MetricDataPoints struct{}"
	}

	return strings.Join([]string{"MetricDataPoints", string(data)}, " ")
}
