package model

import (
	"encoding/json"

	"strings"
)

//
type AgencyTokenUserDomain struct {
	// 委托方A的账号ID。
	Id string `json:"id"`
	// 委托方A的账号名称。
	Name string `json:"name"`
}

func (o AgencyTokenUserDomain) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AgencyTokenUserDomain struct{}"
	}

	return strings.Join([]string{"AgencyTokenUserDomain", string(data)}, " ")
}
