package model

import (
	"encoding/json"

	"strings"
)

// 批量删除伸缩配置请求
type BatchDeleteScalingConfigsRequestBody struct {
	// 伸缩配置ID。
	ScalingConfigurationId *[]string `json:"scaling_configuration_id,omitempty"`
}

func (o BatchDeleteScalingConfigsRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchDeleteScalingConfigsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteScalingConfigsRequestBody", string(data)}, " ")
}
