package model

import (
	"encoding/json"

	"strings"
)

type PeerAddress struct {
	// 域名地址
	DomainPort *string `json:"domain_port,omitempty"`
	// IP地址
	IpPort *string `json:"ip_port,omitempty"`
}

func (o PeerAddress) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PeerAddress struct{}"
	}

	return strings.Join([]string{"PeerAddress", string(data)}, " ")
}
