package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateMemberResponse struct {
	Member         *MemberResp `json:"member,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o CreateMemberResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateMemberResponse struct{}"
	}

	return strings.Join([]string{"CreateMemberResponse", string(data)}, " ")
}
