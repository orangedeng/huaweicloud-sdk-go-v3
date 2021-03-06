package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneListUsersForGroupByAdminRequest struct {
	GroupId string `json:"group_id"`
}

func (o KeystoneListUsersForGroupByAdminRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneListUsersForGroupByAdminRequest struct{}"
	}

	return strings.Join([]string{"KeystoneListUsersForGroupByAdminRequest", string(data)}, " ")
}
