package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type GlanceShowImageMemberRequest struct {
	ImageId  string `json:"image_id"`
	MemberId string `json:"member_id"`
}

func (o GlanceShowImageMemberRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GlanceShowImageMemberRequest struct{}"
	}

	return strings.Join([]string{"GlanceShowImageMemberRequest", string(data)}, " ")
}
