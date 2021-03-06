package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateMfaDeviceResponse struct {
	VirtualMfaDevice *CreateMfaDeviceRespon `json:"virtual_mfa_device,omitempty"`
	HttpStatusCode   int                    `json:"-"`
}

func (o CreateMfaDeviceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateMfaDeviceResponse struct{}"
	}

	return strings.Join([]string{"CreateMfaDeviceResponse", string(data)}, " ")
}
