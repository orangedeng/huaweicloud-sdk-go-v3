package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowMemberRequest struct {
	PoolId   string `json:"pool_id"`
	MemberId string `json:"member_id"`
}

func (o ShowMemberRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowMemberRequest struct{}"
	}

	return strings.Join([]string{"ShowMemberRequest", string(data)}, " ")
}
