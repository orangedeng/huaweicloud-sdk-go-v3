package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneListAuthProjectsResponse struct {
	Links *LinksSelf `json:"links,omitempty"`
	// 项目信息列表。
	Projects       *[]AuthProjectResult `json:"projects,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o KeystoneListAuthProjectsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneListAuthProjectsResponse struct{}"
	}

	return strings.Join([]string{"KeystoneListAuthProjectsResponse", string(data)}, " ")
}
