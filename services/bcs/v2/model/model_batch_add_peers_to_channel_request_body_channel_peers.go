package model

import (
	"encoding/json"

	"strings"
)

type BatchAddPeersToChannelRequestBodyChannelPeers struct {
	// peer加入的通道名称
	ChannelName string `json:"channel_name"`
	// 加入通道peer名称和数量，key为组织名称，value为peer数量
	Peers map[string]int32 `json:"peers"`
}

func (o BatchAddPeersToChannelRequestBodyChannelPeers) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchAddPeersToChannelRequestBodyChannelPeers struct{}"
	}

	return strings.Join([]string{"BatchAddPeersToChannelRequestBodyChannelPeers", string(data)}, " ")
}
