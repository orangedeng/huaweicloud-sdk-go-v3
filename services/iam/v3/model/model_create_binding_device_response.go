package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateBindingDeviceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateBindingDeviceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateBindingDeviceResponse struct{}"
	}

	return strings.Join([]string{"CreateBindingDeviceResponse", string(data)}, " ")
}
