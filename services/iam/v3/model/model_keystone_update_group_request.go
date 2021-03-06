package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneUpdateGroupRequest struct {
	GroupId string                          `json:"group_id"`
	Body    *KeystoneUpdateGroupRequestBody `json:"body,omitempty"`
}

func (o KeystoneUpdateGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneUpdateGroupRequest struct{}"
	}

	return strings.Join([]string{"KeystoneUpdateGroupRequest", string(data)}, " ")
}
