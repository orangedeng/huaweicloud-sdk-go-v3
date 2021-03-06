package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowMemberRequest struct {
	MemberId string `json:"member_id"`
	PoolId   string `json:"pool_id"`
}

func (o ShowMemberRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowMemberRequest struct{}"
	}

	return strings.Join([]string{"ShowMemberRequest", string(data)}, " ")
}
