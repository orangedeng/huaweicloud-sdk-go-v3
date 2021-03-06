package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneRemoveUserFromGroupRequest struct {
	GroupId string `json:"group_id"`
	UserId  string `json:"user_id"`
}

func (o KeystoneRemoveUserFromGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneRemoveUserFromGroupRequest struct{}"
	}

	return strings.Join([]string{"KeystoneRemoveUserFromGroupRequest", string(data)}, " ")
}
