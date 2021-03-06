package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateUserResponse struct {
	User           *CreateUserResult `json:"user,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o CreateUserResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateUserResponse struct{}"
	}

	return strings.Join([]string{"CreateUserResponse", string(data)}, " ")
}
