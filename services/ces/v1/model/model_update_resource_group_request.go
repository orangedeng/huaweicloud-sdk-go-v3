package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateResourceGroupRequest struct {
	ContentType string                          `json:"Content-Type"`
	GroupId     string                          `json:"group_id"`
	Body        *UpdateResourceGroupRequestBody `json:"body,omitempty"`
}

func (o UpdateResourceGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateResourceGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateResourceGroupRequest", string(data)}, " ")
}
