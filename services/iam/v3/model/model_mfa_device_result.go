package model

import (
	"encoding/json"

	"strings"
)

//
type MfaDeviceResult struct {
	// 虚拟MFA的设备序列号。
	SerialNumber string `json:"serial_number"`
	// IAM用户ID。
	UserId string `json:"user_id"`
}

func (o MfaDeviceResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MfaDeviceResult struct{}"
	}

	return strings.Join([]string{"MfaDeviceResult", string(data)}, " ")
}
