package model

import (
	"encoding/json"

	"strings"
)

//
type CreateLoginTokenRequestBody struct {
	Auth *LoginTokenAuth `json:"auth"`
}

func (o CreateLoginTokenRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateLoginTokenRequestBody struct{}"
	}

	return strings.Join([]string{"CreateLoginTokenRequestBody", string(data)}, " ")
}
