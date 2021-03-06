package model

import (
	"encoding/json"

	"strings"
)

//
type ScopedToken struct {
	// 联邦unscoped token的ID。
	Id string `json:"id"`
}

func (o ScopedToken) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ScopedToken struct{}"
	}

	return strings.Join([]string{"ScopedToken", string(data)}, " ")
}
