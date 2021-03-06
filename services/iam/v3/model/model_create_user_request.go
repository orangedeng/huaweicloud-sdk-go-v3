package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateUserRequest struct {
	Body *CreateUserRequestBody `json:"body,omitempty"`
}

func (o CreateUserRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateUserRequest struct{}"
	}

	return strings.Join([]string{"CreateUserRequest", string(data)}, " ")
}
