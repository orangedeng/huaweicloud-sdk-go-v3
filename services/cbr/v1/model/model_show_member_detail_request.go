package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowMemberDetailRequest struct {
	BackupId string `json:"backup_id"`
	MemberId string `json:"member_id"`
}

func (o ShowMemberDetailRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowMemberDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowMemberDetailRequest", string(data)}, " ")
}
