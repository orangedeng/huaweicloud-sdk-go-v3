package model

import (
	"encoding/json"

	"strings"
)

//
type UpdateUserRequestBody struct {
	User *UpdateUserOption `json:"user"`
}

func (o UpdateUserRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateUserRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateUserRequestBody", string(data)}, " ")
}
