package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneListProjectsResponse struct {
	Links *Links `json:"links,omitempty"`
	// 项目信息列表。
	Projects       *[]ProjectResult `json:"projects,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o KeystoneListProjectsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneListProjectsResponse struct{}"
	}

	return strings.Join([]string{"KeystoneListProjectsResponse", string(data)}, " ")
}
