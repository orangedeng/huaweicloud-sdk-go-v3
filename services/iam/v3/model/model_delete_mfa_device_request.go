package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteMfaDeviceRequest struct {
	UserId       string `json:"user_id"`
	SerialNumber string `json:"serial_number"`
}

func (o DeleteMfaDeviceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteMfaDeviceRequest struct{}"
	}

	return strings.Join([]string{"DeleteMfaDeviceRequest", string(data)}, " ")
}
