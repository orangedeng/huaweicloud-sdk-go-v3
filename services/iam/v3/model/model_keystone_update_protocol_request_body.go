package model

import (
	"encoding/json"

	"strings"
)

//
type KeystoneUpdateProtocolRequestBody struct {
	Protocol *ProtocolOption `json:"protocol"`
}

func (o KeystoneUpdateProtocolRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneUpdateProtocolRequestBody struct{}"
	}

	return strings.Join([]string{"KeystoneUpdateProtocolRequestBody", string(data)}, " ")
}
