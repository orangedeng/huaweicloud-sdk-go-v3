package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListResourceGroupRequest struct {
	ContentType string  `json:"Content-Type"`
	GroupName   *string `json:"group_name,omitempty"`
	GroupId     *string `json:"group_id,omitempty"`
	Status      *string `json:"status,omitempty"`
	Start       *int32  `json:"start,omitempty"`
	Limit       *int32  `json:"limit,omitempty"`
}

func (o ListResourceGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListResourceGroupRequest struct{}"
	}

	return strings.Join([]string{"ListResourceGroupRequest", string(data)}, " ")
}
