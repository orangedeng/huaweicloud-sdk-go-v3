package model

import (
	"encoding/json"

	"strings"
)

// 创建项目的请求体。
type KeystoneCreateProjectRequestBody struct {
	Project *KeystoneCreateProjectOption `json:"project"`
}

func (o KeystoneCreateProjectRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneCreateProjectRequestBody struct{}"
	}

	return strings.Join([]string{"KeystoneCreateProjectRequestBody", string(data)}, " ")
}
