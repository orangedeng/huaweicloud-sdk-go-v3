package model

import (
	"encoding/json"

	"strings"
)

type CfgRequestBody struct {
	// 链代码名称
	ChaincodeName string `json:"chaincode_name"`
	// SDK配置文件存放路径
	CertPath string `json:"cert_path"`
	// 通道名称
	ChannelName string `json:"channel_name"`
	// key：组织名，value：该组织下需要下载的peer节点信息
	PeerOrgs map[string][]string `json:"peer_orgs"`
	// key：联盟成员名称，value：该联盟成员peer组织名称hash值数组
	UnionInfo map[string][]string `json:"union_info,omitempty"`
}

func (o CfgRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CfgRequestBody struct{}"
	}

	return strings.Join([]string{"CfgRequestBody", string(data)}, " ")
}
