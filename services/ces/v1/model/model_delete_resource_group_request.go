package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteResourceGroupRequest struct {
	ContentType string `json:"Content-Type"`
	GroupId     string `json:"group_id"`
}

func (o DeleteResourceGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteResourceGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteResourceGroupRequest", string(data)}, " ")
}
