package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateMemberStatusRequest struct {
	MemberId string        `json:"member_id"`
	BackupId string        `json:"backup_id"`
	Body     *UpdateMember `json:"body,omitempty"`
}

func (o UpdateMemberStatusRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateMemberStatusRequest struct{}"
	}

	return strings.Join([]string{"UpdateMemberStatusRequest", string(data)}, " ")
}
