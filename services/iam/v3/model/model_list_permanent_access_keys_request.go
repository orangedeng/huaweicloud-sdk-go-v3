package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListPermanentAccessKeysRequest struct {
	UserId *string `json:"user_id,omitempty"`
}

func (o ListPermanentAccessKeysRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListPermanentAccessKeysRequest struct{}"
	}

	return strings.Join([]string{"ListPermanentAccessKeysRequest", string(data)}, " ")
}
