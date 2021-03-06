package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneUpdateGroupResponse struct {
	Group          *KeystoneGroupResultWithLinksSelf `json:"group,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o KeystoneUpdateGroupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneUpdateGroupResponse struct{}"
	}

	return strings.Join([]string{"KeystoneUpdateGroupResponse", string(data)}, " ")
}
