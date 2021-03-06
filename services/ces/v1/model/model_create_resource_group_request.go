package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateResourceGroupRequest struct {
	ContentType string                          `json:"Content-Type"`
	Body        *CreateResourceGroupRequestBody `json:"body,omitempty"`
}

func (o CreateResourceGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateResourceGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateResourceGroupRequest", string(data)}, " ")
}
