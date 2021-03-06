package model

import (
	"encoding/json"

	"strings"
)

//
type AgencyTokenDomain struct {
	// 委托方A的账号名称。
	Name string `json:"name"`
	// 委托方A的账号ID。
	Id string `json:"id"`
}

func (o AgencyTokenDomain) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AgencyTokenDomain struct{}"
	}

	return strings.Join([]string{"AgencyTokenDomain", string(data)}, " ")
}
