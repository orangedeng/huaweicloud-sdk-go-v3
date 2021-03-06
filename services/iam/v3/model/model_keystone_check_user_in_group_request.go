package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneCheckUserInGroupRequest struct {
	GroupId string `json:"group_id"`
	UserId  string `json:"user_id"`
}

func (o KeystoneCheckUserInGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneCheckUserInGroupRequest struct{}"
	}

	return strings.Join([]string{"KeystoneCheckUserInGroupRequest", string(data)}, " ")
}
