package model

import (
	"encoding/json"

	"strings"
)

//
type PwdPasswordUser struct {
	Domain *PwdPasswordUserDomain `json:"domain"`
	// IAM用户名。
	Name string `json:"name"`
	// IAM用户的登录密码。
	Password string `json:"password"`
}

func (o PwdPasswordUser) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PwdPasswordUser struct{}"
	}

	return strings.Join([]string{"PwdPasswordUser", string(data)}, " ")
}
