package model

import (
	"encoding/json"

	"strings"
)

//
type UnbindMfaDevice struct {
	// 待解绑MFA设备的IAM用户ID。
	UserId string `json:"user_id"`
	// 验证码。
	AuthenticationCode *string `json:"authentication_code,omitempty"`
	// MFA设备序列号。
	SerialNumber string `json:"serial_number"`
}

func (o UnbindMfaDevice) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UnbindMfaDevice struct{}"
	}

	return strings.Join([]string{"UnbindMfaDevice", string(data)}, " ")
}
