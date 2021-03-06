package model

import (
	"encoding/json"

	"strings"
)

//
type UpdateUserInformationRequestBody struct {
	User *UpdateUserInformationOption `json:"user"`
}

func (o UpdateUserInformationRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateUserInformationRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateUserInformationRequestBody", string(data)}, " ")
}
