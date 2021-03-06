package model

import (
	"encoding/json"

	"strings"
)

//
type MappingResult struct {
	// 映射ID。
	Id    string     `json:"id"`
	Links *LinksSelf `json:"links"`
	// 将联邦用户映射为本地用户的规则列表。
	Rules []MappingRules `json:"rules"`
}

func (o MappingResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MappingResult struct{}"
	}

	return strings.Join([]string{"MappingResult", string(data)}, " ")
}
