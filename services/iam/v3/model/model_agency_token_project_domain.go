package model

import (
	"encoding/json"

	"strings"
)

//
type AgencyTokenProjectDomain struct {
	// 委托方A的账号名称。
	Name string `json:"name"`
	// 委托方A的账号ID。
	Id string `json:"id"`
}

func (o AgencyTokenProjectDomain) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AgencyTokenProjectDomain struct{}"
	}

	return strings.Join([]string{"AgencyTokenProjectDomain", string(data)}, " ")
}
