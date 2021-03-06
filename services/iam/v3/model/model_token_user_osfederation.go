package model

import (
	"encoding/json"

	"strings"
)

//
type TokenUserOsfederation struct {
	// 用户组信息列表。
	Groups           []OsfederationGroups          `json:"groups"`
	IdentityProvider *OsfederationIdentityprovider `json:"identity_provider"`
	Protocol         *OsfederationProtocol         `json:"protocol"`
}

func (o TokenUserOsfederation) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TokenUserOsfederation struct{}"
	}

	return strings.Join([]string{"TokenUserOsfederation", string(data)}, " ")
}
