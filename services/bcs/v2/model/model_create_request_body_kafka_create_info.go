package model

import (
	"encoding/json"

	"strings"
)

// kafka实例创建信息
type CreateRequestBodyKafkaCreateInfo struct {
	// kafka实例规格，可选：mini：基准带宽100MB/s，small：基准带宽300MB/s，middle：基准带宽600MB/s，high：基准带宽1200MB/s
	Spec *string `json:"spec,omitempty"`
	// 存储空间(单位：GB），至多9000，mini版至少300，small至少1200，middle至少2400，high至少4800
	Storage *int32 `json:"storage,omitempty"`
	// kafka实例可用区
	Az *string `json:"az,omitempty"`
}

func (o CreateRequestBodyKafkaCreateInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateRequestBodyKafkaCreateInfo struct{}"
	}

	return strings.Join([]string{"CreateRequestBodyKafkaCreateInfo", string(data)}, " ")
}
