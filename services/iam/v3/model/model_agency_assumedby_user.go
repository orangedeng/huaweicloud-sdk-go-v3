package model

import (
	"encoding/json"

	"strings"
)

//
type AgencyAssumedbyUser struct {
	// 被委托方B中IAM用户的用户名。
	Name string `json:"name"`
	// 被委托方B中IAM用户的用户ID。
	Id     string                     `json:"id"`
	Domain *AgencyAssumedbyUserDomain `json:"domain"`
	// 被委托方B中IAM用户的密码过期时间（UTC时间），“”表示密码不过期。
	PasswordExpiresAt string `json:"password_expires_at"`
}

func (o AgencyAssumedbyUser) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AgencyAssumedbyUser struct{}"
	}

	return strings.Join([]string{"AgencyAssumedbyUser", string(data)}, " ")
}
