package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreatePostPaidInstanceRequest struct {
	ProjectId string             `json:"project_id"`
	Body      *CreateInstanceReq `json:"body,omitempty"`
}

func (o CreatePostPaidInstanceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreatePostPaidInstanceRequest struct{}"
	}

	return strings.Join([]string{"CreatePostPaidInstanceRequest", string(data)}, " ")
}
