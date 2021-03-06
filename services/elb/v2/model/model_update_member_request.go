package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateMemberRequest struct {
	MemberId string                   `json:"member_id"`
	PoolId   string                   `json:"pool_id"`
	Body     *UpdateMemberRequestBody `json:"body,omitempty"`
}

func (o UpdateMemberRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateMemberRequest struct{}"
	}

	return strings.Join([]string{"UpdateMemberRequest", string(data)}, " ")
}
