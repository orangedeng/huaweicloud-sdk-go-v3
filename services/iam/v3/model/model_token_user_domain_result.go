package model

import (
	"encoding/json"

	"strings"
)

//
type TokenUserDomainResult struct {
	// IAM用户所属账号名称。
	Name string `json:"name"`
	// IAM用户所属账号ID。
	Id string `json:"id"`
}

func (o TokenUserDomainResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TokenUserDomainResult struct{}"
	}

	return strings.Join([]string{"TokenUserDomainResult", string(data)}, " ")
}
