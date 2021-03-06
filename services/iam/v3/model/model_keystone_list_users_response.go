package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneListUsersResponse struct {
	Links *Links `json:"links,omitempty"`
	// IAM用户信息列表。
	Users          *[]KeystoneListUsersResult `json:"users,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o KeystoneListUsersResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneListUsersResponse struct{}"
	}

	return strings.Join([]string{"KeystoneListUsersResponse", string(data)}, " ")
}
