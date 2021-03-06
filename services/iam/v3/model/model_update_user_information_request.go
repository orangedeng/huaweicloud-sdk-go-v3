package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateUserInformationRequest struct {
	UserId string                            `json:"user_id"`
	Body   *UpdateUserInformationRequestBody `json:"body,omitempty"`
}

func (o UpdateUserInformationRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateUserInformationRequest struct{}"
	}

	return strings.Join([]string{"UpdateUserInformationRequest", string(data)}, " ")
}
