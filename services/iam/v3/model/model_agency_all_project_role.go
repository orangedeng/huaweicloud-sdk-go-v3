package model

import (
	"encoding/json"

	"strings"
)

//
type AgencyAllProjectRole struct {
	// 权限ID。
	Id    string     `json:"id"`
	Links *LinksSelf `json:"links"`
	// 权限名。
	Name string `json:"name"`
}

func (o AgencyAllProjectRole) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AgencyAllProjectRole struct{}"
	}

	return strings.Join([]string{"AgencyAllProjectRole", string(data)}, " ")
}
