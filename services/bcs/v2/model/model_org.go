package model

import (
	"encoding/json"

	"strings"
)

type Org struct {
	// 组织MSP标识
	MspId *string `json:"msp_id,omitempty"`
	// 组织域名
	Domain *string `json:"domain,omitempty"`
	// key:节点名称，value：节点详细信息
	Peers map[string]Node `json:"peers,omitempty"`
}

func (o Org) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Org struct{}"
	}

	return strings.Join([]string{"Org", string(data)}, " ")
}
