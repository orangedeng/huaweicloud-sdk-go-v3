package model

import (
	"encoding/json"

	"strings"
)

//
type IdentityproviderOption struct {
	// 身份提供商描述信息。
	Description *string `json:"description,omitempty"`
	// 身份提供商是否启用，true为启用，false为停用，默认为false。
	Enabled *bool `json:"enabled,omitempty"`
}

func (o IdentityproviderOption) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "IdentityproviderOption struct{}"
	}

	return strings.Join([]string{"IdentityproviderOption", string(data)}, " ")
}
