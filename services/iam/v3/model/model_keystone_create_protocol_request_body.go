package model

import (
	"encoding/json"

	"strings"
)

//
type KeystoneCreateProtocolRequestBody struct {
	Protocol *ProtocolOption `json:"protocol"`
}

func (o KeystoneCreateProtocolRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneCreateProtocolRequestBody struct{}"
	}

	return strings.Join([]string{"KeystoneCreateProtocolRequestBody", string(data)}, " ")
}
