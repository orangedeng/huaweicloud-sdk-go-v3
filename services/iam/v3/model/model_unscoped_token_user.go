package model

import (
	"encoding/json"

	"strings"
)

//
type UnscopedTokenUser struct {
	Domain *TokenDomainResult `json:"domain"`
	// 用户ID。
	Id string `json:"id"`
	// 用户名称。
	Name         string                 `json:"name"`
	OsFederation *TokenUserOsfederation `json:"OS-FEDERATION"`
}

func (o UnscopedTokenUser) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UnscopedTokenUser struct{}"
	}

	return strings.Join([]string{"UnscopedTokenUser", string(data)}, " ")
}
