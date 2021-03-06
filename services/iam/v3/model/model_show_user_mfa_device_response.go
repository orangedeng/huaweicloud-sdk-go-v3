package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowUserMfaDeviceResponse struct {
	VirtualMfaDevice *MfaDeviceResult `json:"virtual_mfa_device,omitempty"`
	HttpStatusCode   int              `json:"-"`
}

func (o ShowUserMfaDeviceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowUserMfaDeviceResponse struct{}"
	}

	return strings.Join([]string{"ShowUserMfaDeviceResponse", string(data)}, " ")
}
