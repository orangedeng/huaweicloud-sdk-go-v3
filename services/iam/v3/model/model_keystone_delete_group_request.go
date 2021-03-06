package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneDeleteGroupRequest struct {
	GroupId string `json:"group_id"`
}

func (o KeystoneDeleteGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneDeleteGroupRequest struct{}"
	}

	return strings.Join([]string{"KeystoneDeleteGroupRequest", string(data)}, " ")
}
