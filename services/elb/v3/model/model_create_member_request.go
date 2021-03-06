package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateMemberRequest struct {
	PoolId string                   `json:"pool_id"`
	Body   *CreateMemberRequestBody `json:"body,omitempty"`
}

func (o CreateMemberRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateMemberRequest struct{}"
	}

	return strings.Join([]string{"CreateMemberRequest", string(data)}, " ")
}
