package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteMemberRequest struct {
	PoolId   string `json:"pool_id"`
	MemberId string `json:"member_id"`
}

func (o DeleteMemberRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteMemberRequest struct{}"
	}

	return strings.Join([]string{"DeleteMemberRequest", string(data)}, " ")
}
