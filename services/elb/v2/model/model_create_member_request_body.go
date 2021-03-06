package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type CreateMemberRequestBody struct {
	Member *CreateMemberReq `json:"member"`
}

func (o CreateMemberRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateMemberRequestBody struct{}"
	}

	return strings.Join([]string{"CreateMemberRequestBody", string(data)}, " ")
}
